package security

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IDExtractor(c *gin.Context) (string, error) {

	authHeader := c.GetHeader("Authorization")
	claims := jwt.MapClaims{}

	jwt.ParseWithClaims(authHeader[7:], claims, func(token *jwt.Token) (interface{}, error) {

		return []byte("secret"), nil
	})

	return claims["sub"].(string), nil
}
