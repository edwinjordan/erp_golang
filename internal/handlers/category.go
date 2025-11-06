package handlers

import (
	"net/http"

	"github.com/edwinjordan/erp_golang/internal/database"
	"github.com/edwinjordan/erp_golang/internal/models"
	"github.com/gin-gonic/gin"
)

type CategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := database.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func GetCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, category)
}

func CreateCategory(c *gin.Context) {
	var req CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := database.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, category)
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	var req CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.Name = req.Name
	category.Description = req.Description

	if err := database.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Category{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
