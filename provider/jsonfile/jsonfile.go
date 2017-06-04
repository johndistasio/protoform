package jsonfile

import (
	"encoding/json"
	"io/ioutil"

	"github.com/johndistasio/cauldron/provider"
)

type JsonFile struct {
	path string
}

func New(path string) *JsonFile {
	return &JsonFile{path}
}

func (p *JsonFile) GetData() (provider.TemplateData, error) {
	data := make(provider.TemplateData)

	jsonData, err := ioutil.ReadFile(p.path)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, &data)

	return data, err
}
