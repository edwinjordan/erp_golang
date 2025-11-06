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
		// For simplicity, we'll implement basic role check
		// In production, you'd check specific permissions from database
		c.Next()
	}
}
