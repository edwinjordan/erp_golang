# ERP Golang

A comprehensive Enterprise Resource Planning (ERP) system built with Go (Golang), featuring authentication, RBAC (Role-Based Access Control), master data management, and Point of Sale (POS) functionality.

## Features

- **Authentication**: User registration and login with JWT tokens
- **RBAC**: Role-Based Access Control for managing user permissions
- **Master Data Management**:
  - Products (Produk)
  - Categories (Kategori)
  - Units (Satuan)
- **POS (Point of Sale)**: Sales transaction management with inventory updates

## Tech Stack

- **Language**: Go 1.21+
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt

## Project Structure

```
erp_golang/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── database/
│   │   └── database.go          # Database connection
│   ├── handlers/
│   │   ├── auth.go              # Authentication handlers
│   │   ├── category.go          # Category handlers
│   │   ├── unit.go              # Unit handlers
│   │   ├── product.go           # Product handlers
│   │   └── pos.go               # POS/Sales handlers
│   ├── middleware/
│   │   └── auth.go              # Authentication middleware
│   └── models/
│       └── models.go            # Database models
└── pkg/
    └── utils/
        └── jwt.go               # JWT utilities
```

## Installation

1. Clone the repository:
```bash
git clone https://github.com/edwinjordan/erp_golang.git
cd erp_golang
```

2. Install dependencies:
```bash
go mod download
```

3. Set up PostgreSQL database:
```bash
createdb erp_db
```

4. Create a `.env` file based on `.env.example`:
```bash
cp .env.example .env
```

5. Update the `.env` file with your database credentials.

## Running the Application

```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:8080` (or the port specified in your `.env` file).

## API Endpoints

### Authentication

#### Register
```http
POST /register
Content-Type: application/json

{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "password123",
  "role_id": 2
}
```

#### Login
```http
POST /login
Content-Type: application/json

{
  "username": "john_doe",
  "password": "password123"
}
```

### Categories (Protected - Requires Authentication)

#### Get All Categories
```http
GET /api/categories
Authorization: Bearer <token>
```

#### Create Category
```http
POST /api/categories
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Electronics",
  "description": "Electronic devices"
}
```

#### Update Category
```http
PUT /api/categories/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Updated Category",
  "description": "Updated description"
}
```

#### Delete Category
```http
DELETE /api/categories/:id
Authorization: Bearer <token>
```

### Units (Protected - Requires Authentication)

#### Get All Units
```http
GET /api/units
Authorization: Bearer <token>
```

#### Create Unit
```http
POST /api/units
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Piece",
  "description": "Individual item"
}
```

### Products (Protected - Requires Authentication)

#### Get All Products
```http
GET /api/products
Authorization: Bearer <token>
```

#### Create Product
```http
POST /api/products
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Laptop",
  "description": "High-performance laptop",
  "category_id": 1,
  "unit_id": 1,
  "price": 1500.00,
  "stock": 10
}
```

#### Update Product
```http
PUT /api/products/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Updated Laptop",
  "description": "Updated description",
  "category_id": 1,
  "unit_id": 1,
  "price": 1600.00,
  "stock": 15
}
```

### POS/Sales (Protected - Requires Authentication)

#### Get All Sales
```http
GET /api/sales
Authorization: Bearer <token>
```

#### Create Sale
```http
POST /api/sales
Authorization: Bearer <token>
Content-Type: application/json

{
  "items": [
    {
      "product_id": 1,
      "quantity": 2
    },
    {
      "product_id": 2,
      "quantity": 1
    }
  ]
}
```

## Default Credentials

After running the application for the first time, you can log in with:
- **Username**: admin
- **Password**: admin123

## Database Models

### User
- ID, Username, Email, Password (hashed), RoleID, Role

### Role
- ID, Name, Description, Permissions

### Permission
- ID, Name, Description

### Category
- ID, Name, Description

### Unit
- ID, Name, Description

### Product
- ID, Name, Description, CategoryID, Category, UnitID, Unit, Price, Stock

### Sale
- ID, UserID, User, Total, SaleItems

### SaleItem
- ID, SaleID, ProductID, Product, Quantity, Price, Subtotal

## Features Overview

1. **Authentication & Authorization**
   - JWT-based authentication
   - Password hashing with bcrypt
   - Role-based access control (RBAC)

2. **Master Data Management**
   - CRUD operations for Categories
   - CRUD operations for Units
   - CRUD operations for Products

3. **Point of Sale (POS)**
   - Create sales transactions
   - Automatic inventory management
   - Transaction tracking with user information

## License

MIT License