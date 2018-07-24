package data

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var jsonTests = map[*bytes.Buffer]Data{
	bytes.NewBufferString("{\"foo\" : \"bar\"}"): {"foo": "bar"},
}

func TestJsonParsing(t *testing.T) {
	for reader, expected := range jsonTests {
		actual, err := NewJson(reader).GetData()
		assert.Nil(t, err)
		assert.Equal(t, actual, expected)
	}
}
