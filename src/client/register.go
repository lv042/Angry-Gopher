package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func readRegisterToken() (string, error) {
	token, err := os.ReadFile("../.token.txt")
	if err != nil {
		log.Fatal("Error reading register token file: ", err)
	}
	return string(token), nil
}
