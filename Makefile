# Go Builder Kit v2 - Makefile
# This Makefile provides common development tasks

.PHONY: all build test clean fmt vet lint coverage install-tools run help

# Variables
GO_VERSION := $(shell go version | cut -d' ' -f3)
BINARY_NAME := builder-gen
BINARY_PATH := ./cmd/builder-gen
PACKAGE := github.com/adil-faiyaz98/go-builder-kit
LDFLAGS := -ldflags "-s -w"

# Default target
all: fmt vet test build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	go build $(LDFLAGS) -o bin/$(BINARY_NAME) $(BINARY_PATH)

# Run tests
test:
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...

# Run tests with coverage
coverage: test
	@echo "Generating coverage report..."
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Run go vet
vet:
	@echo "Running go vet..."
	go vet ./...

# Run linter
lint: install-tools
	@echo "Running linter..."
	staticcheck ./...
	golint ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f bin/$(BINARY_NAME)
	rm -f coverage.out coverage.html
	go clean -cache

# Install development tools
install-tools:
	@echo "Installing development tools..."
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/lint/golint@latest

# Run the binary
run: build
	./bin/$(BINARY_NAME) -help

# Install the binary
install:
	@echo "Installing $(BINARY_NAME)..."
	go install $(LDFLAGS) $(BINARY_PATH)

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	go mod tidy

# Update dependencies
update:
	@echo "Updating dependencies..."
	go get -u ./...
	go mod tidy

# Run security check
security:
	@echo "Running security check..."
	@which govulncheck > /dev/null || (echo "Installing govulncheck..." && go install golang.org/x/vuln/cmd/govulncheck@latest)
	govulncheck ./...

# Generate builders for examples
generate:
	@echo "Generating builders..."
	go run $(BINARY_PATH) -input models -output builders -models-package $(PACKAGE)/models

# Run full CI pipeline locally
ci: install-tools fmt vet lint test security build
	@echo "CI pipeline completed successfully!"

# Release check - run before creating a release
release-check: ci
	@echo "Release check completed!"
	@echo "Ready for release!"

# Help
help:
	@echo "Available targets:"
	@echo "  all          - Run fmt, vet, test, and build"
	@echo "  build        - Build the binary"
	@echo "  test         - Run tests"
	@echo "  coverage     - Run tests with coverage report"
	@echo "  fmt          - Format code"
	@echo "  vet          - Run go vet"
	@echo "  lint         - Run linter"
	@echo "  clean        - Clean build artifacts"
	@echo "  install-tools- Install development tools"
	@echo "  run          - Build and run the binary"
	@echo "  install      - Install the binary"
	@echo "  tidy         - Tidy dependencies"
	@echo "  update       - Update dependencies"
	@echo "  security     - Run security check"
	@echo "  generate     - Generate builders for examples"
	@echo "  ci           - Run full CI pipeline"
	@echo "  release-check- Run pre-release checks"
	@echo "  help         - Show this help"
