GOARCH := $(shell go env GOARCH)
GOOS := $(shell go env GOOS)
VERSION := $(shell git describe --abbrev=0)

.PHONY: build
build:
	go build --ldflags "-w -s"

.PHONY: clean
clean:
	rm linker *.tgz

.PHONY: archive
archive:
	tar -czf ${VERSION}-${GOOS}_${GOARCH}.tgz linker
