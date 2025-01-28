# Variables
APP_NAME := lsw
SRC_DIR := .
BUILD_DIR := ./bin
GO_FILES := $(wildcard $(SRC_DIR)/*.go)

# Default target
all: build

# Build the application
build: lint $(GO_FILES)
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)
	@echo "Build complete: $(BUILD_DIR)/$(APP_NAME)"

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Cleanup complete."

# Run the application
run: build
	@echo "Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

# Test the application
test:
	@echo "Running tests..."
	@go test ./...

# Lint the code
lint:
	@echo "Linting..."
	@golangci-lint run

# Help message
help:
	@echo "Usage:"
	@echo "  make          - Build the application"
	@echo "  make build    - Build the application"
	@echo "  make clean    - Remove build artifacts"
	@echo "  make run      - Run the application"
	@echo "  make test     - Run tests"
	@echo "  make lint     - Run linting checks"

.PHONY: all build clean run test lint help
