package middleware

import (
	"net/http"
	"strings"

	"github.com/edwinjordan/erp_golang/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("roleID", claims.RoleID)
		c.Next()
	}
}

func RBACMiddleware(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID, exists := c.Get("roleID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in context"})
			c.Abort()
			return
		}

		// Role-based access: Admin (roleID: 1) has full access
		// Users (roleID: 2) have limited access
		if roleID.(uint) == 1 {
			// Admin has access to everything
			c.Next()
			return
		}

		// For non-admin users, check specific permissions
		// In a production system, you would check permissions from database
		// For now, we allow basic read operations for all authenticated users
		if requiredPermission == "read" {
			c.Next()
			return
		}

		// Deny access for write operations to non-admin users
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		c.Abort()
	}
}
