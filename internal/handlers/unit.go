package handlers

import (
	"net/http"

	"github.com/edwinjordan/erp_golang/internal/database"
	"github.com/edwinjordan/erp_golang/internal/models"
	"github.com/gin-gonic/gin"
)

type UnitRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func GetUnits(c *gin.Context) {
	var units []models.Unit
	if err := database.DB.Find(&units).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch units"})
		return
	}
	c.JSON(http.StatusOK, units)
}

func GetUnit(c *gin.Context) {
	id := c.Param("id")
	var unit models.Unit
	if err := database.DB.First(&unit, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unit not found"})
		return
	}
	c.JSON(http.StatusOK, unit)
}

func CreateUnit(c *gin.Context) {
	var req UnitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	unit := models.Unit{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := database.DB.Create(&unit).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create unit"})
		return
	}

	c.JSON(http.StatusCreated, unit)
}

func UpdateUnit(c *gin.Context) {
	id := c.Param("id")
	var unit models.Unit
	if err := database.DB.First(&unit, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unit not found"})
		return
	}

	var req UnitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	unit.Name = req.Name
	unit.Description = req.Description

	if err := database.DB.Save(&unit).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update unit"})
		return
	}

	c.JSON(http.StatusOK, unit)
}

func DeleteUnit(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Unit{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete unit"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Unit deleted successfully"})
}
