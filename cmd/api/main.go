package main

import (
	"log"

	"github.com/edwinjordan/erp_golang/internal/config"
	"github.com/edwinjordan/erp_golang/internal/database"
	"github.com/edwinjordan/erp_golang/internal/handlers"
	"github.com/edwinjordan/erp_golang/internal/middleware"
	"github.com/edwinjordan/erp_golang/internal/models"
	"github.com/edwinjordan/erp_golang/pkg/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize JWT
	utils.InitJWT(cfg.JWTSecret)

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate database
	if err := database.DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.Category{},
		&models.Unit{},
		&models.Product{},
		&models.Sale{},
		&models.SaleItem{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Seed initial data
	seedData()

	// Setup router
	router := gin.Default()

	// Public routes
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	// Protected routes
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// Category routes
		categories := api.Group("/categories")
		{
			categories.GET("", handlers.GetCategories)
			categories.GET("/:id", handlers.GetCategory)
			categories.POST("", handlers.CreateCategory)
			categories.PUT("/:id", handlers.UpdateCategory)
			categories.DELETE("/:id", handlers.DeleteCategory)
		}

		// Unit routes
		units := api.Group("/units")
		{
			units.GET("", handlers.GetUnits)
			units.GET("/:id", handlers.GetUnit)
			units.POST("", handlers.CreateUnit)
			units.PUT("/:id", handlers.UpdateUnit)
			units.DELETE("/:id", handlers.DeleteUnit)
		}

		// Product routes
		products := api.Group("/products")
		{
			products.GET("", handlers.GetProducts)
			products.GET("/:id", handlers.GetProduct)
			products.POST("", handlers.CreateProduct)
			products.PUT("/:id", handlers.UpdateProduct)
			products.DELETE("/:id", handlers.DeleteProduct)
		}

		// POS/Sales routes
		sales := api.Group("/sales")
		{
			sales.GET("", handlers.GetSales)
			sales.GET("/:id", handlers.GetSale)
			sales.POST("", handlers.CreateSale)
		}
	}

	// Start server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func seedData() {
	// Create default roles
	var adminRole models.Role
	if err := database.DB.Where("name = ?", "admin").First(&adminRole).Error; err != nil {
		adminRole = models.Role{
			Name:        "admin",
			Description: "Administrator role with full access",
		}
		database.DB.Create(&adminRole)
	}

	var userRole models.Role
	if err := database.DB.Where("name = ?", "user").First(&userRole).Error; err != nil {
		userRole = models.Role{
			Name:        "user",
			Description: "Regular user role",
		}
		database.DB.Create(&userRole)
	}

	// Create default admin user
	var adminUser models.User
	if err := database.DB.Where("username = ?", "admin").First(&adminUser).Error; err != nil {
		adminUser = models.User{
			Username: "admin",
			Email:    "admin@example.com",
			RoleID:   adminRole.ID,
		}
		adminUser.HashPassword("admin123")
		database.DB.Create(&adminUser)
		log.Println("Default admin user created: username=admin, password=admin123")
	}

	// Create sample categories
	categories := []models.Category{
		{Name: "Electronics", Description: "Electronic devices and accessories"},
		{Name: "Food", Description: "Food and beverages"},
		{Name: "Clothing", Description: "Apparel and fashion items"},
	}
	for _, cat := range categories {
		var existing models.Category
		if err := database.DB.Where("name = ?", cat.Name).First(&existing).Error; err != nil {
			database.DB.Create(&cat)
		}
	}

	// Create sample units
	units := []models.Unit{
		{Name: "Piece", Description: "Individual item"},
		{Name: "Kilogram", Description: "Weight in kg"},
		{Name: "Liter", Description: "Volume in liters"},
	}
	for _, unit := range units {
		var existing models.Unit
		if err := database.DB.Where("name = ?", unit.Name).First(&existing).Error; err != nil {
			database.DB.Create(&unit)
		}
	}

	log.Println("Database seeded successfully")
}
