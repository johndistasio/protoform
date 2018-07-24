package data

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	cliTestMap   = map[string]interface{}{"bar": float64(1)}
	cliTestArray = []interface{}{"a", "b", "c"}
	cliTests     = map[string]Data{
		"foo=bar hello=world":                  {"foo": "bar", "hello": "world"},
		" x foo=bar 12345 hello=world abc123 ": {"foo": "bar", "hello": "world"},
		"foo={\"bar\":1} hello=world":          {"foo": cliTestMap, "hello": "world"},
		"foo=[\"a\",\"b\",\"c\"]":              {"foo": cliTestArray},
	}
)

func TestCommandLineParsing(t *testing.T) {
	for cli, expected := range cliTests {
		actual, err := NewCommandLine(strings.Split(cli, " ")).GetData()

		assert.Nil(t, err)
		assert.Equal(t, actual, expected)
	}
}
