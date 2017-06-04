package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/johndistasio/cauldron/provider"
	"github.com/johndistasio/cauldron/provider/commandline"
	"github.com/johndistasio/cauldron/provider/jsonfile"
	"github.com/johndistasio/cauldron/version"

	"github.com/Masterminds/sprig"
)

func init() {
	flag.Usage = func() {
		fmt.Print(`Usage: cauldron [arguments] [template parameters]

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
    -template:
        Path to the template to be rendered. This argument is required.
    -version:
        Print version and build details, then exit.

Template Parameters:
    Template parameters take the form of key=value and are used to populate the
    template. The parameter 'color=red' would be referenced in the template as
    {{ .color }}.

Template:
    A template file must specified with the -template flag. The template must
	use the normal Go text template format.

Example:
    $ cauldron -template car.tmpl color=red kind=sedan > car
`)
	}
}

func renderTemplate(path string, data provider.TemplateData) ([]byte, error) {
	p := filepath.Base(path)
	t := template.New(p)
	t, err := t.Funcs(sprig.TxtFuncMap()).ParseFiles(path)

	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)
	err = t.Execute(b, data)

	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
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
	templatePtr := flag.String("template", "", "")
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

	if len(*templatePtr) == 0 {
		quit(errors.New("no template specified"))
	}

	var data provider.TemplateData
	var err error

	switch {
	case len(*jsonPtr) != 0:
		data, err = jsonfile.New(*jsonPtr).GetData()
	default:
		data, err = commandline.New(flag.Args()).GetData()
	}

	if err != nil {
		quit(err)
	}

	tmpl, err := renderTemplate(*templatePtr, data)

	if err != nil {
		quit(err)
	}

	var file *os.File
	defer file.Close()

	switch {
	case *inplacePtr:
		file, err = os.OpenFile(*templatePtr, os.O_WRONLY|os.O_TRUNC, 0600)
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
