package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	helper "grahamkatana/book-libray/utils"
)

// Auth validates token and authorizes users
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", 1)
		if clientToken == "" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthenticated"})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)

		c.Next()

	}
}
