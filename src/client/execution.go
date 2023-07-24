package main

import (
	"os/exec"
	"strings"
	"time"
)

type CommandResult struct {
	Message      string    `json:"message"`
	ID           int8      `json:"id"`
	TimeOpened   time.Time `json:"time_opened"`
	TimeExecuted time.Time `json:"time_executed"`
	Dir          string    `json:"dir"`
	Executed     bool      `json:"executed"`
	Tries        int8      `json:"tries"`
}

func runCmd(inputCmd string) CommandResult {
	//split up the command and arguments
	modInput := strings.Split(inputCmd, " ")

	cmd := exec.Command(modInput[0], modInput[1:]...)
	res, _ := cmd.CombinedOutput()
	dir, _ := exec.Command("pwd").CombinedOutput()

	//fmt.Println("Command: ", cmd.ProcessState)

	if cmd.ProcessState == nil {
		res = []byte("Could not run command")
	}

	commandResult := CommandResult{
		Message:  string(res),
		Time:     time.Now(),
		Dir:      string(dir),
		Executed: cmd.ProcessState != nil,
	}

	return commandResult
}
