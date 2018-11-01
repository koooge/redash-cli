BIN ?= redash

default: build

get:
	go get ./...

build: get
	go build -o $(BIN)

.PHONY: get build
