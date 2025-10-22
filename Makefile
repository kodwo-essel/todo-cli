# Makefile for building todo CLI for multiple platforms

BINARY_NAME=todo
DIST_DIR=dist

# Default target: build for all platforms
all: clean linux windows mac

clean:
	@echo "Cleaning previous builds..."
	@rm -rf $(DIST_DIR)
	@mkdir -p $(DIST_DIR)

linux:
	@echo "Building Linux binary..."
	GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-linux-amd64

windows:
	@echo "Building Windows binary..."
	GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-windows-amd64.exe

mac:
	@echo "Building macOS binaries..."
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -o $(DIST_DIR)/$(BINARY_NAME)-darwin-arm64

# Optional: build a single platform
.PHONY: clean all linux windows mac
