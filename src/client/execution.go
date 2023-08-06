package main

import (
	"os/exec"
	"strings"
	"time"
)

func runCmd(result CommandResult) CommandResult {
	if result.Tries >= 3 {
		return result
	}

	// Split up the command and arguments
	modInput := strings.Split(result.Command, " ")

	cmd := exec.Command(modInput[0], modInput[1:]...)
	res, _ := cmd.CombinedOutput()
	dir, _ := exec.Command("pwd").CombinedOutput()

	//fmt.Println("Command: ", cmd.ProcessState)

	if cmd.ProcessState == nil {
		res = []byte("Could not run command")
		//increase tries
		result.Tries++
	} else {
		// Only update fields if the command is executed successfully
		result.Executed = true
		result.TimeExecuted = time.Now()
	}

	// We don't need to update TimeOpened, ID, and Command fields

	// Update the Response, Dir, and Tries fields of the CommandResult
	result.Response = string(res)
	result.Dir = string(dir)

	return result
}

func runInstruction(result InstructionResult) InstructionResult {
	if result.Tries >= 3 {
		return result
	}

	// Split up the command and arguments
	modInput := strings.Split(result.Instruction, " ")

	cmd := exec.Command(modInput[0], modInput[1:]...)
	res, _ := cmd.CombinedOutput()

	//fmt.Println("Command: ", cmd.ProcessState)

	if cmd.ProcessState == nil {
		res = []byte("Could not run command")
		//increase tries
		result.Tries++
	} else {
		// Only update fields if the command is executed successfully
		result.Executed = true
		result.TimeExecuted = time.Now()
	}

	// We don't need to update TimeOpened, ID, and Command fields

	// Update the Response, Dir, and Tries fields of the CommandResult
	result.Response = string(res)

	return result
}
