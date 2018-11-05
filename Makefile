BIN ?= redash
VERSION := 0.0.1

LDFLAGS := -ldflags "-X github.com/koooge/redash-cli/cmd.version=$(VERSION)"

default: build

get:
	go get ./...

build: get
	go build -o $(BIN) $(LDFLAGS)

.PHONY: get build
