package authenticate

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const tokenMaxAge = 48 * time.Hour

var secretKey = []byte("many_days+ago*")

var TokenInvalidErr = errors.New("json web token is invalid")

type TokenClaims struct {
	jwt.MapClaims

	UserID uint `json:"user_id"`
}

func EncodeToken(userID uint) (string, error) {
	claims := TokenClaims{
		UserID: userID,
		MapClaims: jwt.MapClaims{
			"exp": time.Now().Add(tokenMaxAge).Unix(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
}

func DecodeToken(token string) (*TokenClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(_ *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse jwt fail: %w", err)
	}
	if !tokenClaims.Valid {
		return nil, TokenInvalidErr
	}
	claims, right := tokenClaims.Claims.(*TokenClaims)
	if !right {
		return nil, TokenInvalidErr
	}
	return claims, nil
}
