BINARY_NAME=cidr-tool
VERSION=$(shell git describe --tags --always --dirty)

all: build

build:
	go build -ldflags "-X main.version=$(VERSION)" -o $(BINARY_NAME) main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $(BINARY_NAME)-linux main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $(BINARY_NAME)-windows.exe main.go

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $(BINARY_NAME)-darwin main.go

clean:
	rm -f $(BINARY_NAME) $(BINARY_NAME)-linux $(BINARY_NAME)-windows.exe $(BINARY_NAME)-darwin

.PHONY: all build build-linux build-windows build-darwin clean
