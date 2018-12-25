package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/johndistasio/cauldron/data"

	"github.com/Masterminds/sprig"
)

var (
	version   = "next"
	commit    = ""
	httpRegex = regexp.MustCompile("^http(s){0,1}://")
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
        Read template data from the specified JSON file or URL. Command-line
		template parameters are ignored.
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

func renderTemplate(t *template.Template, s data.Source) ([]byte, error) {
	d, err := s.GetData()

	if err != nil {
		err = fmt.Errorf("failed to parse data: %s", err.Error())
		return nil, err
	}

	b := new(bytes.Buffer)

	err = t.Execute(b, d)

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
	execPtr := flag.String("exec", "", "")
	filePtr := flag.String("file", "", "")
	helpPtr := flag.Bool("help", false, "")
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
		versionString := fmt.Sprintf("cauldron version=%s", version)
		if commit != "" {
			versionString += fmt.Sprintf(" commit=%s", commit)
		}
		fmt.Println(versionString)
		os.Exit(0)
	}

	if *templatePtr == "" {
		quit(errors.New("no template specified"))
	}

	// Now we can do actual work

	tmplName := filepath.Base(*templatePtr)
	tmpl := template.New(tmplName)
	tmpl, err := tmpl.Funcs(sprig.TxtFuncMap()).ParseFiles(*templatePtr)

	if err != nil {
		quit(err)
	}

	var src data.Source

	switch {
	case httpRegex.MatchString(*jsonPtr):
		src = data.NewHttp(*jsonPtr, nil)
	case *jsonPtr != "":
		src = data.NewJsonFile(*jsonPtr)
	default:
		src = data.NewCommandLine(flag.Args())
	}

	if err != nil {
		quit(err)
	}

	var file *os.File
	defer file.Close()

	switch {
	case *inplacePtr:
		file, err = os.OpenFile(*templatePtr, os.O_WRONLY|os.O_TRUNC, 0600)
	case *filePtr != "":
		file, err = os.OpenFile(*filePtr, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	default:
		file = os.Stdout
	}

	if err != nil {
		quit(err)
	}

	rendered, err := renderTemplate(tmpl, src)

	if err != nil {
		quit(err)
	}

	_, err = file.Write(rendered)

	if err != nil {
		quit(err)
	}

	if *execPtr != "" {
		cmd := strings.Split(*execPtr, " ")
		err := exec.Command(cmd[0], cmd[1:]...).Run()

		if err != nil {
			err = fmt.Errorf("failed to exec: %s", err.Error())
			quit(err)
		}
	}
}
