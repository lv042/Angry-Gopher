package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func VerifyToken(tokenString string) bool {
	stringSecretKey := os.Getenv("JWT_SECRET")

	if stringSecretKey == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(stringSecretKey), nil
	})

	if err != nil {
		return false
	}

	// Check if the token is valid and not expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		return time.Now().Before(expirationTime)
	}

	return false
}

func GenerateToken(host string, id int, validity time.Duration) (string, error) {
	stringSecretKey := os.Getenv("JWT_SECRET")

	if stringSecretKey == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	if host == "" {
		return "", fmt.Errorf("host cannot be empty")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
	claims["host"] = host
	claims["exp"] = time.Now().Add(validity).Unix()

	tokenString, err := token.SignedString([]byte(stringSecretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
