XK6_VERSION := latest
XK6_BINARY := $(shell which xk6 || echo "")

# Targets
.PHONY: all build run test

all: test run

deps:
	@if [ -z "$(XK6_BINARY)" ]; then \
		echo "Installing xk6..."; \
		go install go.k6.io/xk6/cmd/xk6@$(XK6_VERSION); \
	else \
		echo "xk6 is already installed."; \
	fi

build: deps
	@echo "Building xk6 extension..."
	@xk6 build github.com/inditex/lib-xksixjsonparser=.

run: deps
	@echo "Running example..."
	@xk6 run --vus 1 --duration 10s ./examples/main.js

test: deps
	@echo "Running unit tests..."
	@go clean -testcache && go test ./...