package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// this could be any heavy logic to validate the input you require
		if token != "valid-token" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token invalid"})
			c.Abort()
			return

		}

		c.Next()

	}
}
