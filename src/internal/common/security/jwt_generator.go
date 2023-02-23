package security

import (
	"fmt"
	core "github.com/muhammadqazi/SIS-Backend-Go/src/internal/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenManager interface {
	NewJWT(userID string, role string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	NewPasswordResetJWT(userID string) (string, error)
	ValidatePasswordResetToken(token string) (*jwt.Token, error)
}

type tokenManager struct {
	signingKey              string
	signingKeyResetPassword string
	Algorithm               string
	ResetPasswordAlgorithm  string
}

func NewTokenManager(c core.Config) TokenManager {
	return &tokenManager{
		signingKey:              c.SecretKey,
		signingKeyResetPassword: c.ResetPasswordSecretKey,
		Algorithm:               c.JWTAlgorithm,
		ResetPasswordAlgorithm:  c.ResetPasswordAlgorithm,
	}
}

func (t *tokenManager) NewJWT(userID string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		Subject:   role,
		Id:        userID,
	})

	return token.SignedString([]byte(t.signingKey))
}

func (t *tokenManager) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}

		if t_.Method.Alg() != t.Algorithm {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Method.Alg())
		}

		return []byte(t.signingKey), nil
	})
}

/* Forgot Password JWT */

func (t *tokenManager) NewPasswordResetJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		Id:        userID,
	})

	return token.SignedString([]byte(t.signingKeyResetPassword))
}

func (t *tokenManager) ValidatePasswordResetToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}

		if t_.Method.Alg() != t.ResetPasswordAlgorithm {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Method.Alg())
		}

		return []byte(t.signingKeyResetPassword), nil
	})
}
