VERSION = $$(git describe --tags --always --dirty)-$$(git name-rev --name-only HEAD)
BUILD_LDFLAGS = "-X main.Version=$(VERSION)"
BIN = vip
BUILD_DIR = ./build

all: build

deps:
	go get -d -v -t ./...

build: deps
	mkdir -p $(BUILD_DIR)
	go build -ldflags=$(BUILD_LDFLAGS) -o $(BUILD_DIR)/$(BIN)

fmt:
	gofmt -w .
