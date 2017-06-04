package data

import "os"

type JsonFile struct {
	*Json
}

func NewJsonFile(path string) (*JsonFile, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0600)

	if err != nil {
		return nil, err
	}

	return &JsonFile{&Json{file}}, nil
}
