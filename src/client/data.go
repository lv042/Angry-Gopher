package main

import "time"

type Device struct {
	ID              int        `json:"id"`
	ServerURL       string     `json:"server_url"`
	Token           string     `json:"token"`
	SystemInfo      SystemInfo // Updated: Made it public
	CommandList     []CommandResult
	InstructionList []InstructionResult
}

type SystemInfo struct {
	Hostname     string `json:"hostname"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
}

type InstructionResult struct {
	Instruction  string    `json:"instruction"`
	Response     string    `json:"response"`
	ID           int32     `json:"id"`
	TimeOpened   time.Time `json:"time_opened"`
	TimeExecuted time.Time `json:"time_executed"`
	Dir          string    `json:"dir"`
	Executed     bool      `json:"executed"`
	Tries        int32     `json:"tries"`
}

type CommandResult struct {
	Command      string    `json:"command"`
	Response     string    `json:"response"`
	ID           int32     `json:"id"`
	TimeOpened   time.Time `json:"time_opened"`
	TimeExecuted time.Time `json:"time_executed"`
	Dir          string    `json:"dir"`
	Executed     bool      `json:"executed"`
	Tries        int32     `json:"tries"`
}

func newDevice() *Device {
	return &Device{
		ID:              0,
		ServerURL:       "http://localhost:3000",
		SystemInfo:      getSystemInfo(),
		CommandList:     []CommandResult{},
		InstructionList: []InstructionResult{},
	}
}
