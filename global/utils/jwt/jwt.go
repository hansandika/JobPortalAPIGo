package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/hansandika/go-job-portal-api/global/constant"
)

func GenerateNewJwtToken(data jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	tok, err := token.SignedString([]byte(os.Getenv(constant.JWT_SECRET)))
	if err != nil {
		return "", err
	}
	return tok, nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv(constant.JWT_SECRET)), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check Expiry date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, errors.New("token expired")
		}

		if claims["email"] == "" || claims["role_id"] == "" {
			return nil, errors.New("invalid token")
		}

		return claims, nil
	} else {
		return nil, err
	}
}
