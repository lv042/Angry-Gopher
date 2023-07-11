package main

import (
	dotenv "github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"time"
)

// Define the server URL
var serverURL = "https://localhost:3000/"
var sysInfo *SystemInfo

func main() {
	getSystemInfo()
	err := register(serverURL)
	if err != nil {
		log.Info("Error registering system: ", err)
		time.Sleep(10 * time.Second)
		main()
	}

}

func setupDotenv() {
	err := dotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
