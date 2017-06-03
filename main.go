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

type parameters struct {
	Data         map[string]interface{}
	TemplatePath string
}

func init() {
	flag.Usage = func() {
		fmt.Print(`Usage: cauldron [arguments] [template params] template

Arguments:
    -help:
        Print this text and exit.
    -inplace:
        Write in-place instead of to standard output.
    -json:
        Read template data from the specified JSON file. Command-line parameters
		are ignored.
    -version:
        Print version and build details, then exit.

Template Parameters:
    Template arguments take the form of key=value and are used in the template.

Template:
    The last argument that doesn't start with a "-" or include a "=" is used as
    the path to the template. The template must use the normal Go text template
    format.

Example:
    $ cauldron color=red kind=sedan car.tmpl > car
`)
	}
}

func parseParameters(cli []string) parameters {
	params := parameters{
		Data:         make(map[string]interface{}),
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
				params.Data[key] = val
			} else {
				params.Data[key] = complex
			}
		} else {
			params.TemplatePath = arg
		}
	}

	return params
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

	params := parseParameters(flag.Args())

	if len(params.TemplatePath) == 0 {
		quit(errors.New("no template specified"))
	}

	if len(*jsonPtr) != 0 {
		jsondata, err := ioutil.ReadFile(*jsonPtr)
		err = json.Unmarshal(jsondata, &params.Data)

		if err != nil {
			quit(err)
		}
	}

	templ, err := template.New(filepath.Base(params.TemplatePath)).Funcs(
		sprig.TxtFuncMap()).ParseFiles(params.TemplatePath)

	if err != nil {
		quit(err)
	}

	if *inplacePtr {
		file, err := os.OpenFile(params.TemplatePath, os.O_WRONLY|os.O_TRUNC, 0600)
		defer file.Close()

		if err != nil {
			quit(err)
		}

		buf := new(bytes.Buffer)
		err = templ.Execute(buf, params.Data)

		if err != nil {
			quit(err)
		}

		_, err = file.WriteString(buf.String())

	} else {
		err = templ.Execute(os.Stdout, params.Data)
	}

	if err != nil {
		quit(err)
	}
}
