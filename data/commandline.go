package data

import (
	"encoding/json"
	"strings"
)

type CommandLine struct {
	arguments []string
}

func NewCommandLine(args []string) *CommandLine {
	return &CommandLine{args}
}

func (p *CommandLine) GetData() (Data, error) {
	data := make(Data)

	for _, arg := range p.arguments {
		if idx := strings.Index(arg, "="); idx > -1 {
			key := arg[:idx]
			val := arg[idx+1:]

			var complex interface{}
			err := json.Unmarshal([]byte(val), &complex)

			if err != nil {
				// If we can't parse the input as JSON, treat it as plain text.
				data[key] = val
			} else {
				data[key] = complex
			}
		}
	}

	return data, nil
}
