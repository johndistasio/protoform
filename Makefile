# vi: set ft=make:

.PHONY: default
default: clean fmt lint test build smoketest

.PHONY: build
build:
	@goreleaser release --skip-publish --rm-dist --snapshot

.PHONY: test
test:
	@go mod download
	@go test -v ./...

.PHONY: smoketest
smoketest:
	@./smoketest.sh "dist/$(shell go env GOOS)_$(shell go env GOARCH)/cauldron"

.PHONY: fmt
fmt:
	@files=$$(go fmt ./...); \
	if [ -n "$$files" ]; then \
	  echo "Incorrect formatting on:"; \
	  echo $$files; \
	  exit 1; \
	fi

.PHONY: lint
lint:
	@golint -set_exit_status

.PHONY: clean
clean:
	@rm -rf dist/
