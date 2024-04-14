package application

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func DecodeToken(token string) (*Claims, error) {
	var claims Claims

	_, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		fmt.Println("Error parsing claims")
		return nil, err
	}

	return &claims, nil

}

func ValidateToken(token string) (bool, error) {
	tokenSecret := os.Getenv("JWT_SECRET")

	tokenData, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("not expected sign in method")
		}

		return []byte(tokenSecret), nil

	})

	if err != nil {
		return false, err
	}

	if _, ok := tokenData.Claims.(*Claims); ok && tokenData.Valid {
		return true, nil
	} else {
		return false, nil
	}
}
