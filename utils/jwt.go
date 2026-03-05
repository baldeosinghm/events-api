package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Logic for generation and verifying JWT

const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	// For security reasons, do not expose password (not just here, anywhere)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	// SignedString takes a byte slice
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method.")
		}

		// Secret key needs to be returned as a byte slice
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token.")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("Invalid token!")
	}

	// How to extract email and user id

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid token claims.")
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64)) // Claims stores id as float64, so convert it to int64
	return userId, nil
}
