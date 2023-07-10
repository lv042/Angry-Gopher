package main

import (
	dotenv "github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// Define the server URL
var serverURL = "https://localhost:3000/"
var sysInfo *SystemInfo

func main() {
	runCmd("ls -a")

}

func setupDotenv() {
	err := dotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
