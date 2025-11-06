# API Documentation

## Base URL
```
http://localhost:8080
```

## Authentication

All protected endpoints require a JWT token in the Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

---

## Public Endpoints

### 1. Register User

**POST** `/register`

Create a new user account.

**Request Body:**
```json
{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "password123",
  "role_id": 2
}
```

**Response (201 Created):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com",
    "role_id": 2,
    "role": {
      "id": 2,
      "name": "user",
      "description": "Regular user role"
    },
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 2. Login

**POST** `/login`

Authenticate and get a JWT token.

**Request Body:**
```json
{
  "username": "john_doe",
  "password": "password123"
}
```

**Response (200 OK):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com",
    "role_id": 2,
    "role": {
      "id": 2,
      "name": "user",
      "description": "Regular user role"
    }
  }
}
```

---

## Protected Endpoints

### Categories

#### Get All Categories

**GET** `/api/categories`

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "name": "Electronics",
    "description": "Electronic devices and accessories",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

#### Get Single Category

**GET** `/api/categories/:id`

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Electronics",
  "description": "Electronic devices and accessories",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

#### Create Category

**POST** `/api/categories`

**Request Body:**
```json
{
  "name": "Electronics",
  "description": "Electronic devices and accessories"
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "name": "Electronics",
  "description": "Electronic devices and accessories",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

#### Update Category

**PUT** `/api/categories/:id`

**Request Body:**
```json
{
  "name": "Updated Electronics",
  "description": "Updated description"
}
```

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Updated Electronics",
  "description": "Updated description",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

#### Delete Category

**DELETE** `/api/categories/:id`

**Response (200 OK):**
```json
{
  "message": "Category deleted successfully"
}
```

---

### Units

#### Get All Units

**GET** `/api/units`

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "name": "Piece",
    "description": "Individual item",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

#### Get Single Unit

**GET** `/api/units/:id`

#### Create Unit

**POST** `/api/units`

**Request Body:**
```json
{
  "name": "Piece",
  "description": "Individual item"
}
```

#### Update Unit

**PUT** `/api/units/:id`

**Request Body:**
```json
{
  "name": "Updated Piece",
  "description": "Updated description"
}
```

#### Delete Unit

**DELETE** `/api/units/:id`

---

### Products

#### Get All Products

**GET** `/api/products`

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "name": "Laptop",
    "description": "High-performance laptop",
    "category_id": 1,
    "category": {
      "id": 1,
      "name": "Electronics",
      "description": "Electronic devices"
    },
    "unit_id": 1,
    "unit": {
      "id": 1,
      "name": "Piece",
      "description": "Individual item"
    },
    "price": 1500.00,
    "stock": 10,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

#### Get Single Product

**GET** `/api/products/:id`

#### Create Product

**POST** `/api/products`

**Request Body:**
```json
{
  "name": "Laptop",
  "description": "High-performance laptop",
  "category_id": 1,
  "unit_id": 1,
  "price": 1500.00,
  "stock": 10
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "name": "Laptop",
  "description": "High-performance laptop",
  "category_id": 1,
  "category": {
    "id": 1,
    "name": "Electronics",
    "description": "Electronic devices"
  },
  "unit_id": 1,
  "unit": {
    "id": 1,
    "name": "Piece",
    "description": "Individual item"
  },
  "price": 1500.00,
  "stock": 10,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

#### Update Product

**PUT** `/api/products/:id`

**Request Body:**
```json
{
  "name": "Updated Laptop",
  "description": "Updated description",
  "category_id": 1,
  "unit_id": 1,
  "price": 1600.00,
  "stock": 15
}
```

#### Delete Product

**DELETE** `/api/products/:id`

---

### Sales (POS)

#### Get All Sales

**GET** `/api/sales`

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "user_id": 1,
    "user": {
      "id": 1,
      "username": "john_doe",
      "email": "john@example.com"
    },
    "total": 3100.00,
    "sale_items": [
      {
        "id": 1,
        "sale_id": 1,
        "product_id": 1,
        "product": {
          "id": 1,
          "name": "Laptop",
          "price": 1500.00
        },
        "quantity": 2,
        "price": 1500.00,
        "subtotal": 3000.00
      }
    ],
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

#### Get Single Sale

**GET** `/api/sales/:id`

#### Create Sale

**POST** `/api/sales`

**Request Body:**
```json
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

**Response (201 Created):**
```json
{
  "id": 1,
  "user_id": 1,
  "user": {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com"
  },
  "total": 3100.00,
  "sale_items": [
    {
      "id": 1,
      "sale_id": 1,
      "product_id": 1,
      "product": {
        "id": 1,
        "name": "Laptop",
        "price": 1500.00
      },
      "quantity": 2,
      "price": 1500.00,
      "subtotal": 3000.00
    },
    {
      "id": 2,
      "sale_id": 1,
      "product_id": 2,
      "product": {
        "id": 2,
        "name": "Mouse",
        "price": 100.00
      },
      "quantity": 1,
      "price": 100.00,
      "subtotal": 100.00
    }
  ],
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

---

## Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid request body"
}
```

### 401 Unauthorized
```json
{
  "error": "Invalid token"
}
```

### 404 Not Found
```json
{
  "error": "Resource not found"
}
```

### 500 Internal Server Error
```json
{
  "error": "Internal server error"
}
```

---

## Notes

1. All timestamps are in ISO 8601 format (UTC)
2. The JWT token expires after 24 hours
3. Sale creation automatically updates product stock
4. Default admin credentials: username=`admin`, password=`admin123`
5. Role IDs: 1=admin, 2=user
