package data

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var validJsonTests = map[*bytes.Buffer]Data{
	bytes.NewBufferString("{\"foo\" : \"bar\"}"): {"foo": "bar"},
}

func TestJsonParsing(t *testing.T) {
	for reader, expected := range validJsonTests {
		actual, err := NewJson(reader).GetData()
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	}

	_, err := NewJson(new(bytes.Buffer)).GetData()

	assert.NotNil(t, err)
}

func TestJsonFileReading(t *testing.T) {
	actual, err := NewJsonFile("../testdata/treats.json").GetData()

	expected := Data{
		"icecream": []interface{}{"chocolate", "vanilla", "strawberry"},
		"slushes":  []interface{}{"grape", "watermelon", "strawberry"},
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)

	_, err = NewJsonFile("sodifgjsldkfjal").GetData()

	assert.NotNil(t, err)
}
