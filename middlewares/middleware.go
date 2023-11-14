package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nghiack7/ecq-skillspar-league/pkg/models"
)

func RequiredApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Max-Age", "1000")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		var ak string
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			authHeaderParts := strings.Split(authHeader, " ")
			if len(authHeaderParts) == 2 && authHeaderParts[0] == "Bearer" {
				ak = authHeaderParts[1]
			} else if len(authHeaderParts) == 1 {
				ak = authHeaderParts[0]
			}
		}
		if ak == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API Key not set"})
			return
		}
		u, err := models.GetUserByAPIKey(ak)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API Key"})
			return
		}
		c.Set("user", u)
		c.Set("user_id", u.ID)
		c.Set("api_key", ak)
		c.Next()
	}
}
