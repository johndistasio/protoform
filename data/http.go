package data

import (
	"fmt"
	"net/http"
)

type Http struct {
	url     string
	headers map[string]string
}

func NewHttp(url string, headers map[string]string) *Http {
	return &Http{url, headers}
}

func (p *Http) GetData() (Data, error) {
	req, err := http.NewRequest("GET", p.url, nil)

	if err != nil {
		return nil, err
	}

	for k, v := range p.headers {
		req.Header.Add(k, v)
	}

	res, err := (&http.Client{}).Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received http %s", res.Status)
	}

	return (&Json{res.Body}).GetData()
}
