package utils

import (
	"errors"
	"os"
	"time"

	"github.com/DarioRoman01/photos/models"
	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(username, id, tipe string) (string, error) {
	claims := models.Claims{
		Username: username,
		UserID:   id,
		Type:     tipe,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 2)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func VerifyToken(tokenString string) (*models.Claims, error) {
	claims := new(models.Claims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
