# vi: set ft=make:

NAME=protoform
VERSION=0.5.0

BUILT=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

EXCLUDES=$(addprefix --exclude=,build rpmbuild .git .idea .vagrant)

GOARCH=$(shell go env GOARCH)
GOOS=$(shell go env GOOS)
GOVERSION=$(shell go version | awk '{print $$3}' | tr -d 'go')

GOLDFLAGS='$(addprefix -X main.,Name=$(NAME) Version=$(VERSION) Built=$(BUILT) GoVersion=$(GOVERSION) GoOs=$(GOOS) GoArch=$(GOARCH))'

.PHONY: build
default: build

archive:
	mkdir -p build/
	tar $(EXCLUDES) -czvf build/$(NAME)-$(VERSION).tar.gz .

build:
	mkdir -p build/
	go build -ldflags $(GOLDFLAGS) -v -o build/$(NAME) github.com/johndistasio/$(NAME)

clean:
	rm -rf build/
