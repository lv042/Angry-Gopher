package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func setupRoutes(app *fiber.App) {
	protected := app.Group("/").Use(authMiddleware)

	protected.Post("/register", handleRegister)

	protected.Post("/ping/:id", handlePing)

	protected.Get("/cmd/:id", handleGetCommandList)

	protected.Post("/cmd/:id", handleUpdateCommandResult)

	protected.Get("/ins/:id", handleGetInstructionList)

	protected.Post("/ins/:id", handleUpdateInstructionResult)

	protected.Get("/version", handleCheckVersion)

	protected.Get("/devices", handleGetDevices)

	protected.Post("/add/:id", handleAddCommandsAndInstructions)
}

func serverListen(app *fiber.App) {
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
