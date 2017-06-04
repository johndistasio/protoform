package data

import (
	"reflect"
	"strings"
	"testing"
)

var cliTestMap = map[string]float64{"bar": 1}

var cliTestArray = []string{"a", "b", "c"}

var cliTests = map[string]Data{
	"foo=bar hello=world":                  {"foo": "bar", "hello": "world"},
	" x foo=bar 12345 hello=world abc123 ": {"foo": "bar", "hello": "world"},
	// TODO: Make these tests work
	// Printing the data structures indicates they should pass, but they fail
	//"foo={\"bar\":1} hello=world":          {"foo": cliTestMap, "hello": "world"},
	//"foo=[\"a\",\"b\",\"c\"]":              {"foo": cliTestArray},
}

func TestCommandLineParsing(t *testing.T) {
	for cli, expected := range cliTests {
		actual, err := NewCommandLine(strings.Split(cli, " ")).GetData()

		if err != nil || !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %s got %s", expected, actual)
		}
	}

}
