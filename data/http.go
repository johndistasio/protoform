package data

import (
	"net/http"
)

type Http struct {
	url string
}

func NewHttp(url string) *Http {
	return &Http{url}
}

func (p *Http) GetData() (Data, error) {
	res, err := http.Get(p.url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return (&Json{res.Body}).GetData()
}
