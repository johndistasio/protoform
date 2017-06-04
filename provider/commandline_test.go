package provider

import (
	"reflect"
	"strings"
	"testing"
)

var testMap = map[string]float64{"bar": 1}

var testArray = []string{"a", "b", "c"}

var tests = map[string]TemplateData{
	"foo=bar hello=world":                  {"foo": "bar", "hello": "world"},
	" x foo=bar 12345 hello=world abc123 ": {"foo": "bar", "hello": "world"},
	// TODO: Make these tests work
	// Printing the data structures indicates they should pass, but they fail
	//"foo={\"bar\":1} hello=world":          {"foo": testMap, "hello": "world"},
	//"foo=[\"a\",\"b\",\"c\"]":              {"foo": testArray},
}

func TestCommandLineParsing(t *testing.T) {
	for cli, expected := range tests {
		actual, err := NewCommandLine(strings.Split(cli, " ")).GetData()

		if err != nil || !reflect.DeepEqual(actual, expected) {
			t.Errorf("Failed parsing \"%s\" expected %s got %s", cli, expected, actual)
		}
	}

}
