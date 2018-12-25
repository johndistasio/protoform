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

	actual, err := NewHttp(server.URL, nil).GetData()
	expected := Data{"foo": "bar"}

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestHttpBadConfig(t *testing.T) {
	headers := map[string]string{"Content-Type": "application/json"}
	_, err := NewHttp("", headers).GetData()
	assert.NotNil(t, err)
}

func TestHttpBadResponse(t *testing.T) {
	handler := http.HandlerFunc(func(
		w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	_, err := NewHttp(server.URL, nil).GetData()

	assert.NotNil(t, err)
}

func TestHttpHeaderParsing(t *testing.T) {
	handler := http.HandlerFunc(func(
		w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Test-Header") == "" {
			w.Write([]byte("{\"fail\":true}"))
		} else {
			w.Write([]byte("{\"fail\":false}"))
		}
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	headers := map[string]string{"X-Test-Header": "potato"}
	actual, err := NewHttp(server.URL, headers).GetData()
	expected := Data{"fail": false}

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
