package main

import (
	"github.com/gofiber/fiber/v2"
)

var devices []Device
var app = fiber.New()
var appConfig AppConfig

func main() {
	//must be initialized first
	setupDotenv()

	//setup config
	appConfig = newAppConfig()
	checkForSecret()

	//for production
	displayTestJWT()

	//all tasks that need to be done while the server is running
	updateApplication(app)

	//setup all the routes
	setupRoutes(app)

	//start the server and listen on port 3000
	serverListen(app)
}

func updateApplication(app *fiber.App) {
	updateLastOnline()
	logDevices()

	go func() {
		for {
			//TODO: update this to use a database
		}
	}()
}
