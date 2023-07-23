package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"runtime"
	"time"
)

func logDevices() {
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
						"ID":           command.ID,
						"Message":      command.Message,
						"TimeOpened":   command.TimeOpened,
						"TimeExecuted": command.TimeExecuted,
						"Dir":          command.Dir,
						"Executed":     command.Executed,
					}).Info("Command: ", command.Message)
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
			commandResult := newCommandResult(&devices[i], message, rand.Intn(2) == 1) // Assuming all commands are executed for simplicity

			// Add the CommandResult to the CommandList
			devices[i].CommandList = append(devices[i].CommandList, commandResult)
		}
	}
}
