BINARY_NAME = envx
MAIN_PATH = ./cmd/envx
GOBIN_DIR = $(shell go env GOBIN)
ifeq ($(GOBIN_DIR),)
	GOBIN_DIR = $(shell go env GOPATH)/bin
endif

.DEFAULT_GOAL := help

build:
	go build -o $(BINARY_NAME) $(MAIN_PATH)

install:
	go install $(MAIN_PATH)
	@echo "✔ Installed $(BINARY_NAME) to $(GOBIN_DIR)"
	@echo "Ensure $(GOBIN_DIR) is in your PATH"

run:
	go run $(MAIN_PATH)

clean:
	rm -f $(BINARY_NAME)

fmt:
	go fmt ./...

tidy:
	go mod tidy

test-unit:
	go test ./tests/unit/...

test:
	go test ./...

test-cover:
	go test ./tests/unit/... -cover

test-race:
	go test ./tests/unit/... -race

check-path:
	@echo $$PATH | tr ':' '\n' | grep -q "$(GOBIN_DIR)" \
		&& echo "✔ $(GOBIN_DIR) is in PATH" \
		|| echo "⚠ $(GOBIN_DIR) is NOT in PATH. Add it to your shell config."

help:
	@echo "envx Makefile commands:"
	@echo ""
	@echo "  make build       Build the binary locally"
	@echo "  make install     Install envx to $(GOBIN_DIR)"
	@echo "  make run         Run envx without installing"
	@echo "  make check-path  Verify PATH setup"
	@echo "  make fmt         Format code"
	@echo "  make tidy        Clean dependencies"
