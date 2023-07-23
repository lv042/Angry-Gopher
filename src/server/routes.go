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
		ID:         int8(id),
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

func handleGetInstructionList(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 || id > len(devices) {
		return fiber.ErrBadRequest
	}

	device := devices[id-1]

	// Return the instruction list for the device
	return c.JSON(device.InstructionList)
}

func handleCheckVersion(c *fiber.Ctx) error {
	var request struct {
		Version string `json:"version"`
	}

	if err := c.BodyParser(&request); err != nil {
		return fiber.ErrBadRequest
	}

	// Compare the current version with the newest version
	clientVersion := request.Version

	if clientVersion != version {
		return c.JSON(fiber.Map{"message": "You are not on the newest version"})
		//@TODO: Return the newest version
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
		commandResult := newCommandResult(device, cmd, true)

		// Add the CommandResult to the CommandList
		device.CommandList = append(device.CommandList, commandResult)
	}

	// Append instructions to InstructionList
	device.InstructionList.Instructions = append(device.InstructionList.Instructions, request.Instructions...)

	return c.JSON(fiber.Map{"message": "Added commands and instructions to device"})
}
