package data

import (
	gojson "encoding/json"
	"io"
	"io/ioutil"
	"os"
)

type Json struct {
	reader io.Reader
}

type JsonFile struct {
	*Json
}

func NewJson(r io.Reader) *Json {
	return &Json{r}
}

func NewJsonFile(path string) (*JsonFile, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0600)

	if err != nil {
		return nil, err
	}

	return &JsonFile{&Json{file}}, nil
}

func (p *Json) GetData() (Data, error) {
	data := make(Data)

	jsonData, err := ioutil.ReadAll(p.reader)

	if err != nil {
		return nil, err
	}

	err = gojson.Unmarshal(jsonData, &data)

	return data, err
}
