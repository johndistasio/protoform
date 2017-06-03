# vi: set ft=make:

APP     := protoform
VERSION := 0.6.0
VERSION := $(shell git describe --tags --always 2>/dev/null || echo '$(VERSION)')

BUILT=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
REVISION=$(shell git rev-parse --short HEAD)

GO_ARCH=$(shell go env GOARCH)
GO_OS=$(shell go env GOOS)
GO_VERSION=$(shell go version | awk '{print $$3}' | tr -d 'go')

GO_LDFLAGS=$(addprefix -X main.,Name=$(APP) Version=$(VERSION) Built=$(BUILT) GoVersion=$(GO_VERSION) GoOs=$(GO_OS) GoArch=$(GO_ARCH))

TARBALL_EXCLUDE=$(addprefix --exclude=,build rpmbuild .git .idea .vagrant)

.PHONY: build
default: build

archive:
	mkdir -p build/
	tar $(TARBALL_EXCLUDE) -czvf build/$(APP)-$(VERSION).tar.gz .

build:
	mkdir -p build/
	go build -ldflags '$(GO_LDFLAGS)' -v -o build/$(APP) github.com/johndistasio/$(APP)

clean:
	rm -rf build/
