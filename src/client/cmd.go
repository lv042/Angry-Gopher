package main

import (
	"os/exec"
	"strings"
	"time"
)

type CommandResult struct {
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
	Dir      string    `json:"dir"`
	Executed bool      `json:"executed"`
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
