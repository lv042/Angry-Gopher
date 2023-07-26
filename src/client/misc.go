package main

import (
	dotenv "github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"time"
)

func setupDotenv() {
	err := dotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Info(message, err)

		log.Info("Sleeping for 10 seconds and restarting...")
		time.Sleep(10 * time.Second)
		main()
	}
}

func logCommandResult(result CommandResult) {
	log.WithFields(log.Fields{
		"command":       result.Command,
		"response":      result.Response,
		"id":            result.ID,
		"dir":           result.Dir,
		"executed":      result.Executed,
		"tries":         result.Tries,
		"time opened":   result.TimeOpened,
		"time executed": result.TimeExecuted,
	}).Info("Command result")
}
