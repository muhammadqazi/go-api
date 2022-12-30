package security

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenManager interface {
	NewJWT(userID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type tokenManager struct {
	signingKey string
}

func NewTokenManager(signedKey string) TokenManager {
	return &tokenManager{signingKey: signedKey}
}

func (t *tokenManager) NewJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		Subject:   userID,
	})

	return token.SignedString([]byte(t.signingKey))
}

func (t *tokenManager) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(t.signingKey), nil
	})
}
