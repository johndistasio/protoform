# vi: set ft=make:

PACKAGE = github.com/johndistasio/cauldron

GIT_TAG    = $(shell git describe --tags --always 2>/dev/null)
GO_LDFLAGS = $(addprefix -X main.,version=$(GIT_TAG))

TARBALL_EXCLUDE = $(addprefix --exclude=,build dist rpmbuild .git .idea .vagrant)

.PHONY: test build smoketest

default: clean build

archive:
	@mkdir -p build/
	tar $(TARBALL_EXCLUDE) -czvf build/cauldron-$(GIT_TAG).tar.gz .

build:
	@mkdir -p build/
	go mod download
	CGO_ENABLED=0 go build -ldflags '$(GO_LDFLAGS)' -a -o build/cauldron $(PACKAGE)

test:
	go mod download
	go test -v ./...

smoketest:
	bash ./smoketest.sh

fmt:
	@files=$$(go fmt $(PACKAGE)); \
	if [ -n "$$files" ]; then \
	  echo "Incorrect formatting on:"; \
	  echo $$files; \
	  exit 1; \
	fi

lint:
	golint -set_exit_status

clean:
	@rm -rf build/
