BIN ?= redash
VERSION := 0.0.1

LDFLAGS := -ldflags "-X github.com/koooge/redash-cli/cmd.version=$(VERSION)"

default: build

get:
	go get ./...

build: get
	go build -o build/$(BIN) $(LDFLAGS)

clean:
	rm -rf build dest

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

.PHONY: get build clean cross release-deps release
