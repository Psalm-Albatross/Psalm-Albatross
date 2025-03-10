BINARY_NAME=psalm-cidr-tool
VERSION=$(shell git describe --tags --always --dirty)
BIN_DIR=bin

all: build build-linux build-windows build-darwin

build:
	mkdir -p $(BIN_DIR)
	go build -ldflags "-X main.version=$(VERSION)" -o $(BIN_DIR)/$(BINARY_NAME) main.go

build-linux:
	mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $(BIN_DIR)/$(BINARY_NAME)-linux main.go

build-windows:
	mkdir -p $(BIN_DIR)
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $(BIN_DIR)/$(BINARY_NAME)-windows.exe main.go

build-darwin:
	mkdir -p $(BIN_DIR)
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $(BIN_DIR)/$(BINARY_NAME)-darwin main.go

clean:
	rm -rf $(BIN_DIR)

.PHONY: all build build-linux build-windows build-darwin clean
