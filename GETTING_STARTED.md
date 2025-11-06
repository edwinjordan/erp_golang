# Getting Started Guide

This guide will walk you through setting up and using the ERP Golang application.

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Docker and Docker Compose (optional, for easy PostgreSQL setup)

## Quick Setup

### Option 1: Using the Setup Script (Recommended)

```bash
# Clone the repository
git clone https://github.com/edwinjordan/erp_golang.git
cd erp_golang

# Run the setup script
./setup.sh

# Start the application
make run
```

### Option 2: Manual Setup

1. **Clone the repository**
```bash
git clone https://github.com/edwinjordan/erp_golang.git
cd erp_golang
```

2. **Start PostgreSQL using Docker**
```bash
docker-compose up -d
```

3. **Create environment configuration**
```bash
cp .env.example .env
# Edit .env if needed
```

4. **Install dependencies**
```bash
go mod download
```

5. **Run the application**
```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:8080`.

## First Steps

### 1. Login as Admin

Use the default admin credentials:
- **Username**: `admin`
- **Password**: `admin123`

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123"
  }'
```

You'll receive a response with a JWT token:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "admin",
    "email": "admin@example.com",
    "role_id": 1,
    "role": {
      "id": 1,
      "name": "admin",
      "description": "Administrator role with full access"
    }
  }
}
```

**Save this token** - you'll need it for all subsequent requests.

### 2. View Sample Data

The application automatically creates sample categories and units. View them:

```bash
# Get all categories (replace YOUR_TOKEN with the token from login)
curl -X GET http://localhost:8080/api/categories \
  -H "Authorization: Bearer YOUR_TOKEN"
```

Response:
```json
[
  {
    "id": 1,
    "name": "Electronics",
    "description": "Electronic devices and accessories",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  },
  {
    "id": 2,
    "name": "Food",
    "description": "Food and beverages",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  },
  {
    "id": 3,
    "name": "Clothing",
    "description": "Apparel and fashion items",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

### 3. Create a Product

```bash
curl -X POST http://localhost:8080/api/products \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Dell XPS 15",
    "description": "High-performance laptop",
    "category_id": 1,
    "unit_id": 1,
    "price": 1500.00,
    "stock": 10
  }'
```

### 4. Create a Sale (POS Transaction)

```bash
curl -X POST http://localhost:8080/api/sales \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "items": [
      {
        "product_id": 1,
        "quantity": 2
      }
    ]
  }'
```

This will:
- Create a sales transaction
- Automatically deduct the quantity from product stock
- Calculate the total price
- Record which user made the sale

## Common Workflows

### Register a New User

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "cashier1",
    "email": "cashier1@example.com",
    "password": "securepassword",
    "role_id": 2
  }'
```

### Manage Categories

**Create Category:**
```bash
curl -X POST http://localhost:8080/api/categories \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Books",
    "description": "Books and magazines"
  }'
```

**Update Category:**
```bash
curl -X PUT http://localhost:8080/api/categories/1 \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Consumer Electronics",
    "description": "Updated description"
  }'
```

**Delete Category:**
```bash
curl -X DELETE http://localhost:8080/api/categories/1 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Manage Units

The same pattern applies for units at `/api/units`.

### Manage Products

**Update Product:**
```bash
curl -X PUT http://localhost:8080/api/products/1 \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Dell XPS 15 (Updated)",
    "description": "Updated description",
    "category_id": 1,
    "unit_id": 1,
    "price": 1600.00,
    "stock": 15
  }'
```

### View Sales History

```bash
curl -X GET http://localhost:8080/api/sales \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Using Postman

Import the `postman_collection.json` file into Postman:

1. Open Postman
2. Click "Import"
3. Select the `postman_collection.json` file
4. The collection will be imported with all endpoints configured

Set the variables:
- `base_url`: `http://localhost:8080`
- `token`: Your JWT token from login

## Testing

Run the test suite:

```bash
# Run all tests
make test

# Or run specific tests
go test ./internal/models -v
go test ./pkg/utils -v
```

## Building for Production

```bash
# Build binary
make build

# Run the binary
./bin/erp-api
```

## Troubleshooting

### Database Connection Issues

If you get database connection errors:

1. Ensure PostgreSQL is running:
   ```bash
   docker-compose ps
   ```

2. Check database logs:
   ```bash
   docker-compose logs postgres
   ```

3. Verify your `.env` configuration matches your database setup

### Port Already in Use

If port 8080 is already in use, change the `SERVER_PORT` in your `.env` file:
```
SERVER_PORT=8081
```

## Next Steps

- Read the [API Documentation](API_DOCUMENTATION.md) for detailed endpoint information
- Explore the Postman collection for example requests
- Customize the application for your specific needs
- Add more products, categories, and units
- Set up your POS workflow

## Support

For issues or questions, please open an issue on the GitHub repository.
