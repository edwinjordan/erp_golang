.PHONY: build run clean test deps

# Build the application
build:
	@echo "Building application..."
	@go build -o bin/erp-api cmd/api/main.go

# Run the application
run:
	@echo "Running application..."
	@go run cmd/api/main.go

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

# Run with hot reload (requires air)
dev:
	@echo "Running with hot reload..."
	@air

# Install air for hot reload
install-air:
	@go install github.com/air-verse/air@latest
