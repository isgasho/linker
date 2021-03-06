GOARCH := $(shell go env GOARCH)
GOOS := $(shell go env GOOS)
VERSION := $(shell git describe --tags --abbrev=0)

.PHONY: build
build:
	go build --ldflags "-w -s" -o linker cmd/main.go

.PHONY: test
test:
	go test ./... -v
	rm testdata/target*

.PHONY: clean
clean:
	rm linker *.tgz

.PHONY: archive
archive: build
	tar -czf ${VERSION}-${GOOS}_${GOARCH}.tgz linker
