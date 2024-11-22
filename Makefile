# Project metadata
APP_NAME := gomocku
VERSION := 1.0.0
BUILD_DIR := build
RELEASE_DIR := $(BUILD_DIR)/release
MAIN_FILE := ./cmd/gomocku/main.go

# Go settings
GO ?= go
GOFLAGS := -mod=vendor
LDFLAGS := -X main.version=$(VERSION)

# Define platforms and architectures
PLATFORMS := linux darwin windows
ARCHS := amd64 arm64

all: build

build:
	mkdir -p $(BUILD_DIR)
	$(GO) build -o $(RELEASE_DIR)/$(APP_NAME) -ldflags "$(LDFLAGS)" $(MAIN_FILE);

release:
	mkdir -p $(RELEASE_DIR)

	@for platform in $(PLATFORMS); do \
		for arch in $(ARCHS); do \
			echo "Building $(APP_NAME)-$$platform-$$arch..."; \
			extension=""; \
			if [ "$$platform" = "windows" ]; then \
				extension=".exe"; \
			fi; \
			GOOS=$$platform GOARCH=$$arch $(GO) build -o $(RELEASE_DIR)/$(APP_NAME)-$$platform-$$arch$$extension -ldflags "$(LDFLAGS)" $(MAIN_FILE); \
		done; \
	done

	echo "Generating SHA256 Sum..."
	cd $(RELEASE_DIR) && sha256sum --binary ./* > sha256sum.txt

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
	@echo "  release   Build the binary for all platforms and architectures"
	@echo "  test      Run all tests"
	@echo "  run       Run the application"
	@echo "  clean     Remove build artifacts"
	@echo "  fmt       Format Go source files"
	@echo "  lint      Lint the code (requires golangci-lint)"
	@echo "  vet       Run go vet to check for issues"
	@echo "  deps      Install and tidy dependencies"
	@echo "  help      Show this help"

# Targets
.PHONY: all build test run clean fmt lint vet
