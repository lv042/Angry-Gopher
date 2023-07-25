package main

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func setupRoutes(app *fiber.App) {

	//health check should have no auth
	app.Get("/health", handleHealthCheck)

	protected := app.Group("/").Use(authMiddleware).Use(loggingMiddleware)

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

func loggingMiddleware(c *fiber.Ctx) error {
	log.Info(c.Method(), " ", c.Path())
	log.Info("IP: ", c.IP())
	log.Info("Body: ", string(c.Body()), "\n")

	return c.Next()
}

func serverListen(app *fiber.App) {
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
