package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func newCommandResult(device *Device, cmd string, test bool) Command {
	id := len(device.CommandList) + 1
	rt := Command{
		Command:      cmd,
		Instruction:  false,
		Response:     "Not yet executed",
		ID:           int32(id),
		TimeOpened:   time.Now(),
		TimeExecuted: time.Time{},
		Dir:          "Not yet executed",
		Executed:     false,
		Tries:        0,
	}
	if test {
		rt.TimeOpened = time.Time{}
	}

	return rt
}

func eraseAllLocalData() {
	devices = []Device{}
}

type ServerConfig struct {
	DatabaseString string
	Database       string
	SecretKey      string
	Version        string
}

// Version Constants for config values
const (
	Version = "1.0.1"
)

func newAppConfig() ServerConfig {
	appConfig := ServerConfig{
		DatabaseString: os.Getenv("MONGO_CONNECTION_STRING"),
		Database:       "Devices",
		SecretKey:      os.Getenv("JWT_SECRET"),
		Version:        Version,
	}
	if appConfig.SecretKey == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}
	if appConfig.DatabaseString == "" {
		log.Fatal("MONGO_CONNECTION_STRING environment variable not set")
	}
	return appConfig
}

type SystemInfo struct {
	Hostname     string `json:"hostname"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
}

type Device struct {
	ID          int32      `json:"id"`
	SystemInfo  SystemInfo // Updated: Made it public
	LastOnline  int64
	CommandList []Command
}

type Command struct {
	Command      string    `json:"command"`
	Instruction  bool      `json:"instruction"`
	Response     string    `json:"response"`
	ID           int32     `json:"id"`
	TimeOpened   time.Time `json:"time_opened"`
	TimeExecuted time.Time `json:"time_executed"`
	Dir          string    `json:"dir"`
	Executed     bool      `json:"executed"`
	Tries        int32     `json:"tries"`
}
