package data

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpJsonParsing(t *testing.T) {
	handler := http.HandlerFunc(func(
		w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"foo\":\"bar\"}"))
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	actual, err := NewHttp(server.URL).GetData()

	expected := Data{"foo": "bar"}

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
