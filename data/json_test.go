package data

import (
	"bytes"
	"reflect"
	"testing"
)

var jsonTests = map[*bytes.Buffer]Data{
	bytes.NewBufferString("{\"foo\" : \"bar\"}"): {"foo": "bar"},
}

func TestJsonParsing(t *testing.T) {
	for reader, expected := range jsonTests {
		actual, err := NewJson(reader).GetData()
		if err != nil || !reflect.DeepEqual(actual, expected) {
			t.Errorf("`expected %s got %s", expected, actual)
		}
	}
}
