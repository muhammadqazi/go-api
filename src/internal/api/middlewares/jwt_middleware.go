package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/security"
	"net/http"
	"path"
)

func AuthorizeJWT(jwtService security.TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {

		/*
			"""
			Do not validate the token if the route is in the exceptions list.
			"""
		*/
		route := c.Request.URL.Path
		basePath := path.Base(route)

		exceptions := []string{"login"}

		for _, exc := range exceptions {
			if basePath == exc {
				c.Next()
				return
			}
		}

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
			Extract the subject from the token and pass it to the context.
			"""
		*/

		claims := jwt.MapClaims{}
		sub, err := jwt.ParseWithClaims(authHeader[7:], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		role := sub.Claims.(jwt.MapClaims)["sub"]
		id := sub.Claims.(jwt.MapClaims)["jti"]

		c.Set("role", role)
		c.Set("id", id)

		c.Next()
	}
}
