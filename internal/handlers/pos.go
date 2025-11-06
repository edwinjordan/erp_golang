package handlers

import (
	"net/http"

	"github.com/edwinjordan/erp_golang/internal/database"
	"github.com/edwinjordan/erp_golang/internal/models"
	"github.com/gin-gonic/gin"
)

type SaleItemRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type CreateSaleRequest struct {
	Items []SaleItemRequest `json:"items" binding:"required,min=1"`
}

func GetSales(c *gin.Context) {
	var sales []models.Sale
	if err := database.DB.Preload("User").Preload("SaleItems.Product").Find(&sales).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sales"})
		return
	}
	c.JSON(http.StatusOK, sales)
}

func GetSale(c *gin.Context) {
	id := c.Param("id")
	var sale models.Sale
	if err := database.DB.Preload("User").Preload("SaleItems.Product").First(&sale, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sale not found"})
		return
	}
	c.JSON(http.StatusOK, sale)
}

func CreateSale(c *gin.Context) {
	var req CreateSaleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Start transaction
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var total float64
	var saleItems []models.SaleItem

	// Process each item
	for _, item := range req.Items {
		var product models.Product
		if err := tx.First(&product, item.ProductID).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		// Check stock
		if product.Stock < item.Quantity {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock for product: " + product.Name})
			return
		}

		// Calculate subtotal
		subtotal := product.Price * float64(item.Quantity)
		total += subtotal

		saleItem := models.SaleItem{
			ProductID: product.ID,
			Quantity:  item.Quantity,
			Price:     product.Price,
			Subtotal:  subtotal,
		}
		saleItems = append(saleItems, saleItem)

		// Update stock
		product.Stock -= item.Quantity
		if err := tx.Save(&product).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stock"})
			return
		}
	}

	// Create sale
	sale := models.Sale{
		UserID:    userID.(uint),
		Total:     total,
		SaleItems: saleItems,
	}

	if err := tx.Create(&sale).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sale"})
		return
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	// Load relations
	database.DB.Preload("User").Preload("SaleItems.Product").First(&sale, sale.ID)

	c.JSON(http.StatusCreated, sale)
}
