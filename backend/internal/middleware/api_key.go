package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const API_KEY = "TEST123456"

// cek API Key
func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")

		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "API key missing",
			})
			c.Abort()
			return
		}

		if apiKey != API_KEY {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Invalid API key",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
