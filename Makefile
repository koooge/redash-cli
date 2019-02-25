BIN ?= redash
VERSION := 0.1.1

LDFLAGS := -ldflags "-X github.com/koooge/redash-cli/cmd.version=$(VERSION)"

default: build

get:
	go get ./...

build: get
	go build -o build/$(BIN) $(LDFLAGS)

doc:
	rm -rf ./doc && mkdir -p ./doc
	go run redash.go --doc

lint: get
	go tool vet cmd
	gofmt -e -l `find . -name '*.go'`

test: get lint
	go test -v ./cmd/...

coverage: get
	go test -coverprofile=cover.out -covermode=count ./...
	go tool cover -html=cover.out


cross:
	$(MAKE) build GOOS=linux GOARCH=amd64 BIN=$(BIN)_linux_amd64
	$(MAKE) build GOOS=darwin GOARCH=amd64 BIN=$(BIN)_darwin_amd64
	$(MAKE) build GOOS=windows GOARCH=amd64 BIN=$(BIN)_windows_amd64.exe

release-deps:
	go get -u github.com/tcnksm/ghr

release: clean cross release-deps
	mkdir -p dest
	cp ./build/$(BIN)_linux_amd64 ./dest/$(BIN)
	cd dest; tar zcf $(BIN)_linux_amd64_$(VERSION).tar.gz $(BIN)
	cp ./build/$(BIN)_darwin_amd64 ./dest/$(BIN)
	cd dest; tar zcf $(BIN)_darwin_amd64_$(VERSION).tar.gz $(BIN)
	cp ./build/$(BIN)_windows_amd64.exe ./dest/$(BIN).exe
	cd dest; tar zcf $(BIN)_windows_amd64_$(VERSION).zip $(BIN).exe
	rm ./dest/$(BIN) ./dest/$(BIN).exe
	ghr -delete $(VERSION) ./dest/

clean:
	rm -rf build dest

.PHONY: get build doc lint test cross release-deps release clean
