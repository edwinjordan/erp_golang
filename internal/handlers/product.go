package handlers

import (
	"net/http"

	"github.com/edwinjordan/erp_golang/internal/database"
	"github.com/edwinjordan/erp_golang/internal/models"
	"github.com/gin-gonic/gin"
)

type ProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	CategoryID  uint    `json:"category_id" binding:"required"`
	UnitID      uint    `json:"unit_id" binding:"required"`
	Price       float64 `json:"price" binding:"required,min=0"`
	Stock       int     `json:"stock" binding:"min=0"`
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := database.DB.Preload("Category").Preload("Unit").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.DB.Preload("Category").Preload("Unit").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var req ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		UnitID:      req.UnitID,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create product"})
		return
	}

	// Load relations
	database.DB.Preload("Category").Preload("Unit").First(&product, product.ID)

	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var req ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.Name = req.Name
	product.Description = req.Description
	product.CategoryID = req.CategoryID
	product.UnitID = req.UnitID
	product.Price = req.Price
	product.Stock = req.Stock

	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update product"})
		return
	}

	// Load relations
	database.DB.Preload("Category").Preload("Unit").First(&product, product.ID)

	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
