package main

import (
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

func parseArgs(args []string) Args {
	pargs := newArgs()

	for _, arg := range args {
		if strings.Index(arg, "-") == 0 {
			switch arg {
			case "-h", "-help":
				printUsage(os.Stdout)
				os.Exit(0)
			default:
				fmt.Fprintf(os.Stderr, "Unrecognized argument %s\n", arg)
				printUsage(os.Stderr)
				os.Exit(1)
			}
		} else if idx := strings.Index(arg, "="); idx > -1 {
			pargs.Data[arg[:idx]] = arg[idx+1:]
		} else {
			pargs.Template = arg
		}
	}
	return pargs
}

func main() {
	args := parseArgs(os.Args)

	t, err := template.ParseFiles(args.Template)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse template")
		os.Exit(1)
	}

	err = t.Execute(os.Stdout, args.Data)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to output template")
		os.Exit(1)
	}
}
