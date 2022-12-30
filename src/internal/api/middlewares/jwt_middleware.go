package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	security "github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/security"
)

func AuthorizeJWT(jwtService security.TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {

		/*
			"""
			Get the token from the Authorization header.
			"""
		*/

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Authorization token is required"})
			return
		}
		/*
			"""
			Split the token from the Bearer prefix and validate the token.
			"""
		*/
		_, err := jwtService.ValidateToken(authHeader[7:])

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Invalid authorization token"})
			return
		}

		/*
			"""
			If you want to know what's inside the token, you can use the following code.
			"""
		*/

	}
}
