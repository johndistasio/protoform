package data

import (
	gojson "encoding/json"
	"io"
	"io/ioutil"
)

type Json struct {
	reader io.Reader
}

func NewJson(r io.Reader) *Json {
	return &Json{r}
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
