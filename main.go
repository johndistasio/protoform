package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/johndistasio/cauldron/version"

	"github.com/Masterminds/sprig"
)

type Parameters map[string]interface{}

type Configuration struct {
	TemplateData Parameters
	TemplatePath string
}

func init() {
	flag.Usage = func() {
		fmt.Print(`Usage: cauldron [arguments] [template parameters] template

Arguments:
    -help:
        Print this text and exit.
    -inplace:
        Render the template in-place (overwriting the template) instead of to
        standard output.
    -json:
        Read template data from the specified JSON file. Command-line template
        parameters are ignored.
    -version:
        Print version and build details, then exit.

Template Parameters:
    Template parameters take the form of key=value and are used to populate the
    template. The parameter 'color=red' would be referenced in the template as
    {{ .color }}.

Template:
    The last argument that doesn't start with a "-" or include a "=" is used as
    the path to the template. The template must use the normal Go text template
    format.

Example:
    $ cauldron color=red kind=sedan car.tmpl > car
`)
	}
}

func parseParameters(cli []string) Configuration {
	c := Configuration{
		TemplateData: make(Parameters),
		TemplatePath: "",
	}

	for _, arg := range cli {
		if idx := strings.Index(arg, "="); idx > -1 {
			key := arg[:idx]
			val := arg[idx+1:]

			var complex interface{}
			err := json.Unmarshal([]byte(val), &complex)

			if err != nil {
				// If we can't parse the input as JSON, treat it as plain text.
				c.TemplateData[key] = val
			} else {
				c.TemplateData[key] = complex
			}
		} else {
			c.TemplatePath = arg
		}
	}

	return c
}

func quit(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func main() {
	helpPtr := flag.Bool("help", false, "")
	inplacePtr := flag.Bool("inplace", false, "")
	jsonPtr := flag.String("json", "", "")
	versionPtr := flag.Bool("version", false, "")
	flag.Parse()

	// The flag package doesn't seem to respect long-form "help"?
	if *helpPtr {
		flag.Usage()
		os.Exit(0)
	}

	if *versionPtr {
		fmt.Printf("cauldron %s\n", version.ComputeVersionString())
		os.Exit(0)
	}

	config := parseParameters(flag.Args())

	if len(config.TemplatePath) == 0 {
		quit(errors.New("no template specified"))
	}

	if len(*jsonPtr) != 0 {
		jsonData, err := ioutil.ReadFile(*jsonPtr)
		err = json.Unmarshal(jsonData, &config.TemplateData)

		if err != nil {
			quit(err)
		}
	}

	tmpl, err := template.New(filepath.Base(config.TemplatePath)).Funcs(
		sprig.TxtFuncMap()).ParseFiles(config.TemplatePath)

	if err != nil {
		quit(err)
	}

	if *inplacePtr {
		file, err := os.OpenFile(config.TemplatePath, os.O_WRONLY|os.O_TRUNC, 0600)
		defer file.Close()

		if err != nil {
			quit(err)
		}

		buf := new(bytes.Buffer)
		err = tmpl.Execute(buf, config.TemplateData)

		if err != nil {
			quit(err)
		}

		_, err = file.WriteString(buf.String())

	} else {
		err = tmpl.Execute(os.Stdout, config.TemplateData)
	}

	if err != nil {
		quit(err)
	}
}
