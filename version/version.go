package version

import (
	"fmt"
	"strings"
)

var version_string = "protoform "

var (
	built     string
	goarch    string
	goos      string
	goversion string
	revision  string
	tag       string
	version   string
)

func ComputeVersionString() string {

	if version != "" {
		version_string += fmt.Sprintf("version=%s ", version)
	}

	if tag != "" {
		version_string += fmt.Sprintf("tag=%s ", tag)
	}

	if revision != "" {
		version_string += fmt.Sprintf("revision=%s ", revision)
	}

	if goversion != "" {
		version_string += fmt.Sprintf("goversion=%s ", goversion)
	}

	if goos != "" {
		version_string += fmt.Sprintf("goos=%s ", goos)
	}

	if goarch != "" {
		version_string += fmt.Sprintf("goarch=%s ", goarch)
	}

	if built != "" {
		version_string += fmt.Sprintf("built=%s ", built)
	}

	return strings.TrimSpace(version_string)
}
