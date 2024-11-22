# Project metadata
APP_NAME := gomocku
VERSION := 1.0.0
BUILD_DIR := build
GO_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Go settings
GO ?= go
GOFLAGS := -mod=vendor
LDFLAGS := -X main.version=$(VERSION)

# Targets
.PHONY: all build test run clean fmt lint vet

all: build

build:
	@echo "Building $(APP_NAME)..."
	mkdir -p $(BUILD_DIR)/$(APP_NAME)
	$(GO) build -o $(BUILD_DIR)/$(APP_NAME) -ldflags "$(LDFLAGS)" ./...

test:
	@echo "Running tests..."
	$(GO) test ./...

run: build
	@echo "Running $(APP_NAME)..."
	./$(BUILD_DIR)/$(APP_NAME)

clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

lint:
	@echo "Linting code..."
	golangci-lint run

vet:
	@echo "Running go vet..."
	$(GO) vet ./...

deps:
	@echo "Installing dependencies..."
	$(GO) mod tidy
	$(GO) mod vendor

# Help documentation
help:
	@echo "Usage: make [target]"
	@echo
	@echo "Targets:"
	@echo "  all       Build the project (default target)"
	@echo "  build     Build the binary"
	@echo "  test      Run all tests"
	@echo "  run       Run the application"
	@echo "  clean     Remove build artifacts"
	@echo "  fmt       Format Go source files"
	@echo "  lint      Lint the code (requires golangci-lint)"
	@echo "  vet       Run go vet to check for issues"
	@echo "  deps      Install and tidy dependencies"
	@echo "  help      Show this help"
