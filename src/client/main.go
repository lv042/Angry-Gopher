package main

import (
	log "github.com/sirupsen/logrus"
	"time"
)

// Define the server URL
var device *Device

func main() {
	setupDotenv()
	getSystemInfo()
	device = newDevice()

	var err error
	device.Token, err = readRegisterToken()
	log.Info("Register token: ", device.Token)
	checkError("Error while reading register token: ", err)

	device.ID, err = register(device.ServerURL, device.Token, &device.SystemInfo)
	checkError("Error while registering: ", err)
	time.Sleep(5 * time.Second)

	work()
}

func work() {
	log.Info("Starting to receive tasks")

	commandWorkFlow()
	time.Sleep(99999 * time.Second)
}

func commandWorkFlow() {
	go func() {
		for {
			time.Sleep(1 * time.Second)

			remoteCommands, err := getCommands(device.Token, device.ID)
			checkError("Error while getting commands: ", err)

			if len(remoteCommands) == 0 {
				log.Info("Received no commands")
				continue
			}

			if len(remoteCommands) == len(device.CommandList) {
				log.Info("Received no new commands")
				continue
			}

			device.CommandList = remoteCommands
			// Repeat 3 times
			for i := 1; i <= 3; i++ {
				for idx, command := range device.CommandList {
					if command.Executed {
						continue
					}
					if command.Tries >= 3 {
						log.Infof("Command '%s' has been tried 3 times, but still not executed.", command.Command)
						continue
					}

					updatedCommand := runCmd(command) // Get the updated command

					device.CommandList[idx] = updatedCommand // Update the original slice with the updated command

					logCommandResult(updatedCommand)

					err = postCommandResult(device.Token, device.ID, updatedCommand)
					checkError("Error while posting command result: ", err)

					log.Info("Posted command result")
				}
			}
		}
	}()
}
