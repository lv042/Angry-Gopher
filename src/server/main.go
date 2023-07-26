package main

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var devices []Device
var app = fiber.New()
var appConfig ServerConfig
var mongoClient *mongo.Client
var collection *mongo.Collection
var syncDB = true

func main() {
	setup()

	//all tasks that need to be done while the server is running
	updateApplication(app)

	startServer()
}

func updateApplication(app *fiber.App) {
	updateLastOnline()
	logDevices()
	monitorChanges()
}

func setup() {
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

	//create token for app build
	registerToken, err := GenerateToken("register", 0, time.Hour*24*30)
	checkFatalError("Error generating register token: ", err)

	err = createTokenFile(registerToken)
	checkFatalError("Error creating token file: ", err)

}

func startServer() {
	//setup all the routes
	setupRoutes(app)

	//start the server and listen on port 3000
	serverListen(app)
}
