package provider

import (
	"encoding/json"
	"io/ioutil"
)

type JsonFile struct {
	path string
}

func NewJsonFile(path string) *JsonFile {
	return &JsonFile{path}
}

func (p *JsonFile) GetData() (TemplateData, error) {
	data := make(TemplateData)

	jsonData, err := ioutil.ReadFile(p.path)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, &data)

	return data, err
}
