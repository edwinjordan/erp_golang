package handlers

import (
	"net/http"

	"github.com/edwinjordan/erp_golang/internal/database"
	"github.com/edwinjordan/erp_golang/internal/models"
	"github.com/edwinjordan/erp_golang/pkg/utils"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	RoleID   uint   `json:"role_id"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  models.User  `json:"user"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Default role is user (roleID: 2), admin is 1
	if req.RoleID == 0 {
		req.RoleID = 2
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		RoleID:   req.RoleID,
	}

	if err := user.HashPassword(req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or email already exists"})
		return
	}

	// Load role
	database.DB.Preload("Role").First(&user, user.ID)

	token, err := utils.GenerateToken(user.ID, user.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, AuthResponse{
		Token: token,
		User:  user,
	})
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := user.CheckPassword(req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Load role
	database.DB.Preload("Role").First(&user, user.ID)

	token, err := utils.GenerateToken(user.ID, user.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, AuthResponse{
		Token: token,
		User:  user,
	})
}
