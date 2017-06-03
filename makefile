# vi: set ft=make:

VERSION = 0.6.0
PACKAGE = github.com/johndistasio/cauldron

GIT_REVISION = $(shell git rev-parse --short HEAD)
GIT_TAG      = $(shell git describe --tags --always 2>/dev/null)

GO_ARCH    = $(shell go env GOARCH)
GO_OS      = $(shell go env GOOS)
GO_VERSION = $(shell go version | awk '{print $$3}' | tr -d 'go')
GO_LDFLAGS = $(addprefix -X $(PACKAGE)/version.,version=$(VERSION) revision=$(GIT_REVISION) tag=$(GIT_TAG) goarch=$(GO_ARCH) goos=$(GO_OS) goversion=$(GO_VERSION))

TARBALL_EXCLUDE = $(addprefix --exclude=,build rpmbuild .git .idea .vagrant)

.PHONY: build
default: build

archive:
	@mkdir -p build/
	tar $(TARBALL_EXCLUDE) -czvf build/cauldron-$(VERSION).tar.gz .

build:
	@mkdir -p build/
	go build -ldflags '$(GO_LDFLAGS)' -v -o build/cauldron github.com/johndistasio/cauldron

clean:
	@rm -rf build/
