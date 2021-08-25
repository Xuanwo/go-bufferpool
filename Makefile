SHELL := /bin/bash

.PHONY: check format test

help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  check               to format, vet and lint"
	@echo "  build               to create bin directory and build"
	@echo "  test                to execute unit tests"

check: format vet

format:
	go fmt ./...

vet:
	go vet ./...

build: tidy check
	go build ./...

test:
	go test -race -v ./...

bench:
	go test -v -benchmem -bench=. ./...

tidy:
	go mod tidy && go mod verify
