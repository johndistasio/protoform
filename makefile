# vi: set ft=make:

APP_NAME=protoform
APP_VERSION=0.5.0

BUILT=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

GO_ARCH=$(shell go env GOARCH)
GO_OS=$(shell go env GOOS)
GO_VERSION=$(shell go version | awk '{print $$3}' | tr -d 'go')

GO_LDFLAGS=$(addprefix -X main.,Name=$(APP_NAME) Version=$(APP_VERSION) Built=$(BUILT) GoVersion=$(GO_VERSION) GoOs=$(GO_OS) GoArch=$(GO_ARCH))

TARBALL_EXCLUDE=$(addprefix --exclude=,build rpmbuild .git .idea .vagrant)

.PHONY: build
default: build

archive:
	mkdir -p build/
	tar $(TARBALL_EXCLUDE) -czvf build/$(APP_NAME)-$(APP_VERSION).tar.gz .

build:
	mkdir -p build/
	go build -ldflags '$(GO_LDFLAGS)' -v -o build/$(APP_NAME) github.com/johndistasio/$(APP_NAME)

clean:
	rm -rf build/
