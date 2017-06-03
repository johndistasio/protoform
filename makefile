# vi: set ft=make:

VERSION = 0.6.0
BUILT   = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

GIT_REVISION = $(shell git rev-parse --short HEAD)
GIT_TAG      = $(shell git describe --tags --always 2>/dev/null)

GO_ARCH    = $(shell go env GOARCH)
GO_OS      = $(shell go env GOOS)
GO_VERSION = $(shell go version | awk '{print $$3}' | tr -d 'go')
GO_LDFLAGS = $(addprefix -X main.,Name=protoform Version=$(VERSION) Built=$(BUILT) GitRevision=$(GIT_REVISION) GitTag=$(GIT_TAG) GoArch=$(GO_ARCH) GoOs=$(GO_OS) GoVersion=$(GO_VERSION))

TARBALL_EXCLUDE = $(addprefix --exclude=,build rpmbuild .git .idea .vagrant)

.PHONY: build
default: build

archive:
	mkdir -p build/
	tar $(TARBALL_EXCLUDE) -czvf build/protoform-$(VERSION).tar.gz .

build:
	mkdir -p build/
	go build -ldflags '$(GO_LDFLAGS)' -v -o build/protoform github.com/johndistasio/protoform

clean:
	rm -rf build/
