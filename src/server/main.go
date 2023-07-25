package main

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var devices []Device
var app = fiber.New()
var appConfig AppConfig
var mongoClient *mongo.Client
var collection *mongo.Collection
var syncDB = true

func main() {
	//must be initialized first
	setupDotenv()

	//setup config
	appConfig = newAppConfig()
	checkForSecret()

	// Initialize MongoDB client
	mongoClient = initMongoClient()

	// Get the MongoDB collection
	collection = mongoClient.Database("AngryGopher").Collection("Devices")

	// Load devices data from MongoDB
	loadDevicesFromMongoDB(collection)

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
	monitorChanges()
}
