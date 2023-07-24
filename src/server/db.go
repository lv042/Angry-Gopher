package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func test() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://admin:hwEK2Ys5T7zY87SO@angrygopher.czgt7v8.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func initMongoClient() *mongo.Client {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(appConfig.DatabaseString)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	return client
}

func loadDevicesFromMongoDB() {
	ctx := context.Background()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal("Error fetching devices from MongoDB:", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var device Device
		if err := cursor.Decode(&device); err != nil {
			log.Println("Error decoding device:", err)
			continue
		}
		devices = append(devices, device)
	}
}

func monitorChanges() {
	go func() {

		for {
			length := len(devices)
			if len(devices) == 0 {
				time.Sleep(3 * time.Second)
				continue
			}
			time.Sleep(3 * time.Second)
			if length != len(devices) {
				updateDevicesToMongoDB()
				continue
			}
			time.Sleep(7 * time.Second)
			updateDevicesToMongoDB()
		}

	}()
}

// Function to update changes to MongoDB
func updateDevicesToMongoDB() {
	ctx := context.Background()

	// Step 1: Delete all devices from MongoDB
	_, err := collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		log.Println("Error deleting all devices from MongoDB:", err)
	}

	// Step 2: Insert all devices from the devices array to MongoDB
	documents := make([]interface{}, len(devices))
	for i, device := range devices {
		documents[i] = device
	}

	_, err = collection.InsertMany(ctx, documents)
	if err != nil {
		log.Println("Error inserting devices to MongoDB:", err)
	}
}
