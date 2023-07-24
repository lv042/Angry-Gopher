package main

import (
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"runtime"
	"time"
)

func logDevices() {
	time.Sleep(3 * time.Second) // Wait for the server to start
	go func() {
		for {
			for _, device := range devices {
				log.WithFields(log.Fields{
					"ID":         device.ID,
					"SystemInfo": device.SystemInfo,
					"LastOnline": device.LastOnline,
				}).Info("Logging device ------------------------> ")

				for _, command := range device.CommandList {
					log.WithFields(log.Fields{
						"Command":      command.Command,
						"Response":     command.Response,
						"ID":           command.ID,
						"TimeOpened":   command.TimeOpened,
						"TimeExecuted": command.TimeExecuted,
						"Dir":          command.Dir,
						"Executed":     command.Executed,
						"Tries":        command.Tries,
					}).Info("Command: ", command.Command)
				}
				for _, instruction := range device.InstructionList {
					log.WithFields(log.Fields{
						"Instruction":  instruction.Instruction,
						"Response":     instruction.Response,
						"ID":           instruction.ID,
						"TimeOpened":   instruction.TimeOpened,
						"TimeExecuted": instruction.TimeExecuted,
						"Dir":          instruction.Dir,
						"Executed":     instruction.Executed,
						"Tries":        instruction.Tries,
					}).Info("Instruction: ", instruction.Instruction)
				}
			}

			log.WithFields(log.Fields{
				"NumberOfDevices":    len(devices),
				"NumberOfGoroutines": runtime.NumGoroutine(),
			}).Info("General Information ------------------->")

			time.Sleep(10 * time.Second)
		}
	}()
}

func generateFakeData() {
	devices = []Device{
		{
			ID: 1,
			SystemInfo: SystemInfo{
				Hostname:     "Device1",
				OS:           "Linux",
				Architecture: "x86_64",
			},
		},
		{
			ID: 2,
			SystemInfo: SystemInfo{
				Hostname:     "Device2",
				OS:           "Windows",
				Architecture: "amd64",
			},
		},
	}

	for i := range devices {
		commands := []string{"ls -a", "ls", "pwd"}

		for _, cmd := range commands {
			// Generate a random message and directory for each command
			message := fmt.Sprintf(cmd)

			// Create a new CommandResult instance using the newCommandResult function
			commandResult := newCommandResult(&devices[i], message) // Assuming all commands are executed for simplicity

			// Add the CommandResult to the CommandList
			devices[i].CommandList = append(devices[i].CommandList, commandResult)
		}
	}

	for i := range devices {
		instructions := []string{"install", "update", "uninstall"}

		for _, instruction := range instructions {
			// Generate a random message and directory for each command
			message := fmt.Sprintf(instruction)

			// Create a new InstructionResult instance using the newInstructionResult function
			instructionResult := newInstructionResult(&devices[i], message) // Assuming all instructions are executed for simplicity

			// Add the InstructionResult to the InstructionList
			devices[i].InstructionList = append(devices[i].InstructionList, instructionResult)
		}

	}
}

func updateLastOnline() {
	go func() {
		for {
			for i := range devices {
				// Increment the LastOnline count by 1
				devices[i].LastOnline++
			}

			time.Sleep(1 * time.Second)
		}
	}()
}

func setupDotenv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Warn("Error loading .env file: ", err)
		log.Warn("This is not a problem if you are running the server in production")
	}

}

func checkForSecret() {
	if appConfig.SecretKey == "" {
		log.Info("JWT_SECRET environment variable not set")
	}
}

func displayTestJWT() {
	testJWT, err := GenerateToken("admin", -1, time.Hour*24*30)
	if err != nil {
		log.Warn("Error generating test JWT: ", err)
	}
	//log highlighted
	log.WithFields(log.Fields{
		"JWT": testJWT,
	}).Info("Test JWT")
}
