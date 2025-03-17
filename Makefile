.DEFAULT_GOAL := build

fmt: imports
	go fmt .
.PHONY: fmt

imports:
	goimports -l -w .
.PHONY: imports

build: fmt
	go build -o bin/game main.go
.PHONY: build