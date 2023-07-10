package main

import (
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var stringSecretKey = ""
var secretKey = []byte(stringSecretKey)

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
