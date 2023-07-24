package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func newCommandResult(device *Device, cmd string) CommandResult {
	id := len(device.CommandList) + 1
	return CommandResult{
		Command:      cmd,
		Response:     "Not yet executed",
		ID:           int8(id),
		TimeOpened:   time.Now(),
		TimeExecuted: time.Time{},
		Dir:          "Not yet executed",
		Executed:     false,
		Tries:        0,
	}
}

func newInstructionResult(device *Device, ins string) InstructionResult {
	id := len(device.InstructionList) + 1
	return InstructionResult{
		Instruction:  ins,
		Response:     "Not yet executed",
		ID:           int8(id),
		TimeOpened:   time.Now(),
		TimeExecuted: time.Time{},
		Dir:          "Not yet executed",
		Executed:     false,
		Tries:        0,
	}
}

func eraseAllData() {
	devices = []Device{}
}

type AppConfig struct {
	DatabaseURL string
	Database    string
	SecretKey   string
	Version     string
}

// Constants for config values
const (
	DatabaseURL = ""
	Database    = ""
	Version     = "1.0.0"
)

func newAppConfig() AppConfig {
	appConfig := AppConfig{
		DatabaseURL: DatabaseURL,
		Database:    Database,
		SecretKey:   os.Getenv("JWT_SECRET"),
		Version:     Version,
	}
	if appConfig.SecretKey == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}
	return appConfig
}

type SystemInfo struct {
	Hostname     string `json:"hostname"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
}

type Device struct {
	ID              int8       `json:"id"`
	SystemInfo      SystemInfo // Updated: Made it public
	LastOnline      int64
	CommandList     []CommandResult
	InstructionList []InstructionResult
}

type InstructionResult struct {
	Instruction  string    `json:"instruction"`
	Response     string    `json:"response"`
	ID           int8      `json:"id"`
	TimeOpened   time.Time `json:"time_opened"`
	TimeExecuted time.Time `json:"time_executed"`
	Dir          string    `json:"dir"`
	Executed     bool      `json:"executed"`
	Tries        int8      `json:"tries"`
}

type CommandResult struct {
	Command      string    `json:"command"`
	Response     string    `json:"response"`
	ID           int8      `json:"id"`
	TimeOpened   time.Time `json:"time_opened"`
	TimeExecuted time.Time `json:"time_executed"`
	Dir          string    `json:"dir"`
	Executed     bool      `json:"executed"`
	Tries        int8      `json:"tries"`
}
