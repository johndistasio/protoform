package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
)

type parameters struct {
	Data     map[string]interface{}
	Template string
}

func newParameters() parameters {
	return parameters{
		Data:     make(map[string]interface{}),
		Template: "",
	}
}

func parseParameters(cli []string) parameters {
	params := newParameters()

	for _, arg := range cli {
		if idx := strings.Index(arg, "="); idx > -1 {
			key := arg[:idx]
			val := arg[idx+1:]

			var complex interface{}
			err := json.Unmarshal([]byte(val), &complex)

			if err != nil {
				// If we can't parse the input as JSON, treat it as a plain
				// string.
				params.Data[key] = val
			} else {
				params.Data[key] = complex
			}
		} else {
			params.Template = arg
		}
	}

	return params
}

func init() {
	flag.Usage = func() {
		fmt.Print(`Usage: protoform [args] [template params] template

Arguments:
    -help:
        Print this text and exit.
    -inplace:
        Write in-place instead of to standard output.

Template Parameters:
    Template arguments take the form of key=value and are used in the template.

Template:
    The last argument that doesn't start with a "-" or include a "=" is used as
    the path to the template. The template must use the normal Go text template
    format.

Example:
    $ protoform color=red kind=sedan car.tmpl > car
`)
	}

}

func exitOnError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func main() {
	helpPtr := flag.Bool("help", false, "")
	inplacePtr := flag.Bool("inplace", false, "")
	flag.Parse()

	// The flag package doesn't seem to respect long-form "help"?
	if *helpPtr {
		flag.Usage()
	}

	params := parseParameters(flag.Args())
	t, err := template.New(filepath.Base(params.Template)).Funcs(
		sprig.TxtFuncMap()).ParseFiles(params.Template)

	if err != nil {
		exitOnError(err)
	}

	if *inplacePtr {
		f, err := os.OpenFile(params.Template, os.O_WRONLY|os.O_TRUNC, 0600)
		defer f.Close()

		if err != nil {
			exitOnError(err)
		}

		buf := new(bytes.Buffer)
		err = t.Execute(buf, params.Data)

		if err != nil {
			exitOnError(err)
		}

		_, err = f.WriteString(buf.String())

	} else {
		err = t.Execute(os.Stdout, params.Data)
	}

	if err != nil {
		exitOnError(err)
	}
}
