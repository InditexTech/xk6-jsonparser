XK6_VERSION := v0.13.4
XK6_BINARY := $(shell command -v xk6 2> /dev/null)

GOLANGCI_VERSION := v1.64.5
GOLANGCI_BINARY := $(shell command -v golangci-lint 2> /dev/null)

.PHONY: all
all: format lint test run

.PHONY: deps
deps:
	@if [ -z "$(XK6_BINARY)" ]; then \
		echo "Installing xk6..."; \
		go install go.k6.io/xk6/cmd/xk6@$(XK6_VERSION); \
	else \
		echo "xk6 is already installed."; \
	fi

	@if [ -z "$(GOLANGCI_BINARY)" ]; then \
			echo "Installing golangci-lint..."; \
			go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_VERSION); \
	else \
		echo "golangci-lint is already installed."; \
	fi

.PHONY: build
build: deps
	@echo "Building k6 with jsonparser extension..."
	@xk6 build --with github.com/InditexTech/xk6-jsonparser=.

.PHONY: run
run: deps
	@echo "Running example..."
	@xk6 run ./examples/main.js

.PHONY: verify
verify: format lint test run
	@echo "Running verify..."

.PHONY: test
test:
	@echo "Running unit tests..."
	@go clean -testcache && go test ./...

.PHONY: tidy
tidy:
	@echo "Running go mod tidy..."
	@go mod tidy

.PHONY: format
format:
	@echo "Running go fmt..."
	go fmt ./...

.PHONY: lint
lint: deps
	@echo "Running golangci-lint..."
	@golangci-lint run
