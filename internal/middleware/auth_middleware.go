package middleware

import (
	"log"
	"net/http"

	"gin-be/internal/tool"

	"github.com/gin-gonic/gin"
)

func AnonymousOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are already logged in."})
			return
		}
		c.Next()
	}
}

func AnonymousAndAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := tool.ExtractToken(c)
		if user != nil {
			c.Set("user_id", user.ID.String())
		}
		c.Next()
	}
}

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := tool.ExtractToken(c)
		if err != nil {
			log.Printf("middleware.go|AuthJWT|%s\n", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		// Add user information to the context
		c.Set("user_id", user.ID.String())
		c.Next()
	}
}
