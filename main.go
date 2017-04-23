package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

type Args struct {
	Data     map[string]interface{}
	Template string
}

func newArgs() Args {
	return Args{
		Data:     make(map[string]interface{}),
		Template: "",
	}
}

func printUsage(w io.Writer) {
	fmt.Fprint(w, `Usage: protoform [args] [template args] template

Arguments:
    -h, -help: print this text and exit

Template Arguments:
    Template arguments take the form of key=value and are used in the template.

Template:
    The last argument that doesn't start with a "-" or include a "=" is used as
    the path to the template. The template must use the normal Go text template
    format.

Example:
    protoform color=red kind=sedan car.tmpl > car
`)
}

func parseArgs(args []string) (Args, error) {
	pargs := newArgs()
	var perr error = nil

	for _, arg := range args {
		if strings.Index(arg, "-") == 0 {
			switch arg {
			case "-h", "-help":
				printUsage(os.Stdout)
				os.Exit(0)
			default:
				perr = errors.New(fmt.Sprintf("unrecognized argument %s", arg))
			}

		} else if idx := strings.Index(arg, "="); idx > -1 {
			key := arg[:idx]
			val := arg[idx+1:]

			var complex interface{}
			err := json.Unmarshal([]byte(val), &complex)

			if err != nil {
				// Assume the parser errors on unquoted strings and treat the
				// value as such.
				// TODO: Verify this assumption is valid
				pargs.Data[key] = val
			} else {
				pargs.Data[key] = complex
			}
		} else {
			pargs.Template = arg
		}
	}
	return pargs, perr
}

func main() {
	args, err := parseArgs(os.Args)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse arguments: %s\n", err.Error())
		os.Exit(1)
	}

	t, err := template.ParseFiles(args.Template)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse template: %s\n", err.Error())
		os.Exit(1)
	}

	err = t.Execute(os.Stdout, args.Data)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to render template: %s\n", err.Error())
		os.Exit(1)
	}
}
