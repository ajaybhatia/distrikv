.PHONY: build run test clean profile

BINARY_NAME=distrikvd
CMD_DIR=./cmd/distrikvd

build:
	@echo "Compiling DistriKV production binary..."
	@go build -ldflags="-s -w" -o bin/$(BINARY_NAME) $(CMD_DIR)/main.go

run: build
	@./bin/$(BINARY_NAME)

test:
	@echo "Running all data engine integration tests..."
	@go test -v -race ./...

clean:
	@echo "Cleaning up build artifacts..."
	@rm -rf bin/
	@go clean -cache

profile:
	@echo "Analyzing runtime CPU/Memory profiles..."
	@go tool pprof bin/$(BINARY_NAME) cpu.prof
