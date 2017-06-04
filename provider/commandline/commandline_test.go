package commandline

import (
	"reflect"
	"strings"
	"testing"

	"github.com/johndistasio/cauldron/provider"
)

var simpleTests = map[string]provider.TemplateData{
	"foo=bar hello=world":                  {"foo": "bar", "hello": "world"},
	" x foo=bar 12345 hello=world abc123 ": {"foo": "bar", "hello": "world"},
}

func TestCommandLineParsing(t *testing.T) {
	for cli, expected := range simpleTests {
		actual, err := New(strings.Split(cli, " ")).GetData()

		if err != nil || !reflect.DeepEqual(actual, expected) {
			t.Errorf("Failed parsing \"%s\" expected %s got %s", cli, expected, actual)
		}
	}

}
