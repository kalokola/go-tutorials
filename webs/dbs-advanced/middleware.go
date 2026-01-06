package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates API key and attaches user to context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get API key from header
		apiKey := c.GetHeader("X-API-Key")
		
		// Also check Authorization header as fallback
		if apiKey == "" {
			authHeader := c.GetHeader("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				apiKey = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}

		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, ErrorResponse{
				Success: false,
				Error:   "API key is required. Provide it in X-API-Key header",
			})
			c.Abort()
			return
		}

		// Find user by API key
		var user User
		result := DB.Where("api_key = ?", apiKey).First(&user)
		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, ErrorResponse{
				Success: false,
				Error:   "Invalid API key",
			})
			c.Abort()
			return
		}

		// Attach user to context
		c.Set("user", user)
		c.Set("user_id", user.ID)

		c.Next()
	}
}

// GetCurrentUser retrieves authenticated user from context
func GetCurrentUser(c *gin.Context) (User, bool) {
	user, exists := c.Get("user")
	if !exists {
		return User{}, false
	}
	return user.(User), true
}

// GetCurrentUserID retrieves authenticated user ID from context
func GetCurrentUserID(c *gin.Context) (uint, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}
	return userID.(uint), true
}