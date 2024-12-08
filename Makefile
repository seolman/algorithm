BUILD_DIR := bin
CMD_DIR := cmd/algo

BINARY := algo

run: build
	@echo "Running the application..."
	@$(BUILD_DIR)/$(BINARY)

build:
	@echo "Building the application..."
	@go build -o $(BUILD_DIR)/$(BINARY) $(CMD_DIR)/main.go

test:
	@echo "Running tests..."
	@go test ./...

clean:
	@echo "Cleaning up..."
	@rm -f $(BUILD_DIR)/$(BINARY)

fmt:
	@echo "Formatting the code..."
	@go fmt ./...

lint:
	@echo "Linting the code..."
	@golangci-lint run

install:
	@echo "Installing dependencies..."
	@go mod tidy

help:
	@echo "Usage:"
	@echo "  make build    Build the application"
	@echo "  make run      Run the application"
	@echo "  make test     Run tests"
	@echo "  make clean    Clean up build artifacts"
	@echo "  make fmt      Format the code"
	@echo "  make lint     Lint the code"
	@echo "  make install  Install dependencies"
	@echo "  make help     Show this help message"

.PHONY: build run test clean fmt lint install help
