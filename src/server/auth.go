package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"os"
	"time"
)

var stringSecretKey = ""
var secretKey = []byte(stringSecretKey)

func VerifyToken(tokenString string) bool {
	stringSecretKey = os.Getenv("JWT_SECRET")

	if (stringSecretKey) == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return false
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	} else {
		return false
	}
}

func CreateToken(host string) (string, error) {
	stringSecretKey = os.Getenv("JWT_SECRET")

	if (stringSecretKey) == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = host
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
