package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func handleRegister(c *fiber.Ctx) error {
	var systemInfo SystemInfo

	if err := json.Unmarshal(c.Body(), &systemInfo); err != nil {
		return fiber.ErrBadRequest
	}

	// Get the highest id and increment
	id := len(devices) + 1
	device := Device{
		ID:         int32(id),
		SystemInfo: systemInfo,
		LastOnline: 0,
	}
	devices = append(devices, device)

	// Return their id
	return c.JSON(fiber.Map{"id": id})
}

func handlePing(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 || id > len(devices) {
		return fiber.ErrBadRequest
	}

	device := devices[id-1]
	device.LastOnline = 0
	devices[id-1] = device

	return c.JSON(fiber.Map{"message": "Pong!"})
}

func handleGetCommandList(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 || id > len(devices) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Device with ID " + strconv.Itoa(id) + " does not exist",
		})
	}

	device := devices[id-1]

	// Get the list of CommandResults for the device
	commandResults := make([]CommandResult, len(device.CommandList))
	copy(commandResults, device.CommandList)

	return c.JSON(commandResults)
}

func handleGetInstructionList(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 || id > len(devices) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Device with ID " + strconv.Itoa(id) + " does not exist",
		})
	}

	device := devices[id-1]

	// Get the list of InstructionResults for the device
	instructionResults := make([]InstructionResult, len(device.InstructionList))
	copy(instructionResults, device.InstructionList)

	return c.JSON(instructionResults)
}

func handleUpdateCommandResult(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 || id > len(devices) {
		return fiber.ErrBadRequest
	}

	var updatedCommandResult CommandResult
	if err := c.BodyParser(&updatedCommandResult); err != nil {
		return fiber.ErrBadRequest
	}

	// Find the index of the CommandResult in the device's CommandList
	var commandIndex int
	device := devices[id-1]
	for i, cmdResult := range device.CommandList {
		if cmdResult.ID == updatedCommandResult.ID {
			commandIndex = i
			break
		}
	}

	// Check if the command has not been executed before updating the executed timestamp
	if device.CommandList[commandIndex].Executed {
		return c.JSON(fiber.Map{"message": "Command has already been executed"})
	}

	// Replace the old CommandResult with the updated one
	device.CommandList[commandIndex] = updatedCommandResult
	devices[id-1] = device

	return c.JSON(fiber.Map{"message": "Command result updated successfully"})
}

func handleUpdateInstructionResult(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 || id > len(devices) {
		return fiber.ErrBadRequest
	}

	var updatedInstructionResult InstructionResult
	if err := c.BodyParser(&updatedInstructionResult); err != nil {
		return fiber.ErrBadRequest
	}

	// Find the index of the InstructionResult in the device's InstructionList
	var instructionIndex int
	device := devices[id-1]
	for i, insResult := range device.InstructionList {
		if insResult.ID == updatedInstructionResult.ID {
			instructionIndex = i
			break
		}
	}

	// Check if the instruction has not been executed before updating the executed timestamp
	if device.InstructionList[instructionIndex].Executed {
		return c.JSON(fiber.Map{"message": "Instruction has already been executed"})
	}

	// Replace the old InstructionResult with the updated one
	device.InstructionList[instructionIndex] = updatedInstructionResult
	devices[id-1] = device

	return c.JSON(fiber.Map{"message": "Instruction result updated successfully"})
}

func handleCheckVersion(c *fiber.Ctx) error {
	var request struct {
		Version string `json:"version"`
		Windows bool   `json:"windows"`
	}

	if err := c.BodyParser(&request); err != nil {
		return fiber.ErrBadRequest
	}

	// Compare the current version with the newest version
	clientVersion := request.Version

	if clientVersion != appConfig.Version {
		if request.Windows {

		} else {

		}
	}

	return c.JSON(fiber.Map{"message": "You are on the newest version"})
}

func handleGetDevices(c *fiber.Ctx) error {
	return c.JSON(devices)
}

func handleAddCommandsAndInstructions(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 || id > len(devices) {
		return fiber.ErrBadRequest
	}

	var request struct {
		Commands     []string `json:"commands"`
		Instructions []string `json:"instructions"`
	}

	if err := c.BodyParser(&request); err != nil {
		return fiber.ErrBadRequest
	}

	device := &devices[id-1] // Pass a pointer to Device

	// Generate new CommandResult instances for each command in request.Commands
	for _, cmd := range request.Commands {
		// Create a new CommandResult instance using the newCommandResult function
		commandResult := newCommandResult(device, cmd, false)

		// Add the CommandResult to the CommandList
		device.CommandList = append(device.CommandList, commandResult)
	}

	// Append instructions to InstructionList
	for _, instruction := range request.Instructions {
		instructionResult := newInstructionResult(device, instruction, false)

		// Add the InstructionResult to the InstructionList
		device.InstructionList = append(device.InstructionList, instructionResult)

	}

	return c.JSON(fiber.Map{"message": "Added commands and instructions to device"})
}

func handleHealthCheck(ctx *fiber.Ctx) error {
	//return 200
	return ctx.SendStatus(fiber.StatusOK)
}
