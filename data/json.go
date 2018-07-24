package data

import (
	js "encoding/json"
	"io"
	"io/ioutil"
	"os"
)

type Json struct {
	reader io.Reader
}

type JsonFile struct {
	path string
}

func NewJson(r io.Reader) *Json {
	return &Json{r}
}

func NewJsonFile(path string) *JsonFile {
	return &JsonFile{path}
}

func (p *Json) GetData() (Data, error) {
	data := make(Data)

	jsonData, err := ioutil.ReadAll(p.reader)

	if err != nil {
		return nil, err
	}

	err = js.Unmarshal(jsonData, &data)

	return data, err
}

func (p *JsonFile) GetData() (Data, error) {
	file, err := os.OpenFile(p.path, os.O_RDONLY, 0600)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	return (&Json{file}).GetData()
}
