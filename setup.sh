#!/bin/bash

echo "==================================="
echo "ERP Golang - Quick Start Script"
echo "==================================="
echo ""

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "Docker is not installed. Please install Docker first."
    exit 1
fi

# Check if Docker Compose is installed
if ! command -v docker-compose &> /dev/null; then
    echo "Docker Compose is not installed. Please install Docker Compose first."
    exit 1
fi

echo "Starting PostgreSQL database..."
docker-compose up -d

echo ""
echo "Waiting for database to be ready..."
until docker-compose exec -T postgres pg_isready -U postgres > /dev/null 2>&1; do
    echo "Waiting for PostgreSQL..."
    sleep 2
done
echo "Database is ready!"

echo ""
echo "Creating .env file from .env.example..."
if [ ! -f .env ]; then
    cp .env.example .env
    echo ".env file created successfully"
else
    echo ".env file already exists"
fi

echo ""
echo "Installing Go dependencies..."
go mod download

echo ""
echo "Building application..."
make build

echo ""
echo "==================================="
echo "Setup complete!"
echo "==================================="
echo ""
echo "To start the application, run:"
echo "  make run"
echo ""
echo "Or directly:"
echo "  go run cmd/api/main.go"
echo ""
echo "Default admin credentials:"
echo "  Username: admin"
echo "  Password: admin123"
echo ""
echo "API Documentation: API_DOCUMENTATION.md"
echo "Postman Collection: postman_collection.json"
echo "==================================="
