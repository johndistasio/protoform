package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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
    -exec:
	    Run the specified command after successful template rendering. The
        command does not run in a shell so redirection, pipes, etc. won't work.
    -file:
        Render the template to the specified path instead of to standard output.
    -inplace:
        Render the template in-place (overwriting the template) instead of to
        standard output. Takes precedence over -file.
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

func renderTemplate(c Configuration) ([]byte, error) {
	p := filepath.Base(c.TemplatePath)
	t := template.New(p)
	t, err := t.Funcs(sprig.TxtFuncMap()).ParseFiles(c.TemplatePath)

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, c.TemplateData)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func quit(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func main() {
	helpPtr := flag.Bool("help", false, "")
	execPtr := flag.String("exec", "", "")
	filePtr := flag.String("file", "", "")
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

	tmpl, err := renderTemplate(config)

	if err != nil {
		quit(err)
	}

	var file *os.File
	defer file.Close()

	switch {
	case *inplacePtr:
		file, err = os.OpenFile(config.TemplatePath, os.O_WRONLY|os.O_TRUNC, 0600)
	case len(*filePtr) != 0:
		file, err = os.OpenFile(*filePtr, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	default:
		file = os.Stdout
	}

	if err != nil {
		quit(err)
	}

	_, err = file.Write(tmpl)

	if err != nil {
		quit(err)
	}

	if len(*execPtr) != 0 {
		cmd := strings.Split(*execPtr, " ")
		err := exec.Command(cmd[0], cmd[1:]...).Run()

		if err != nil {
			err = errors.New(fmt.Sprintf("failed to exec: %s", err.Error()))
			quit(err)
		}
	}
}
