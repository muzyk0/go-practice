# Makefile for Wish List Application

.PHONY: help
help: ## Show this help message
	@echo "Wish List Application - Development Commands"
	@echo ""
	@grep -E '^[a-zA-Z_0-9%-]+:.*?## .*$$' $(word 1,$(MAKEFILE_LIST)) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: setup
setup: ## Set up the development environment
	@echo "Setting up development environment..."
	go mod tidy

.PHONY: backend
backend: ## Start the backend server
	@echo "Starting backend server..."
	go run cmd/skill/main.go cmd/skill/flags.go cmd/skill/gzip.go

.PHONY: format
format: ## Format
	@echo "Formatting..."
	go fmt ./...

.PHONY: test
test: ## Run tests
	@echo "Running tests..."
	go test ./...

.PHONY: lint
lint-backend: ## Run golangci-lint
	@echo "Running golangci-lint..."
	golangci-lint run

.PHONY: build
build: ## Build all components
	@echo "Building all components..."
	go build -o bin/server cmd/skill/main.go cmd/skill/flags.go cmd/skill/gzip.go

