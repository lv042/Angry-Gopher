package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func initMongoClient() *mongo.Client {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(appConfig.DatabaseString)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	return client
}

func loadDevicesFromMongoDB(collection *mongo.Collection) {
	eraseAllData()
	ctx := context.Background()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal("Error fetching devices from MongoDB:", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatal("Error closing cursor:", err)
		}
	}(cursor, ctx)

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
			if !syncDB {
				length := len(devices)
				if len(devices) == 0 {
					time.Sleep(3 * time.Second)
					continue
				}
				time.Sleep(3 * time.Second)
				if length != len(devices) {
					updateDevicesToMongoDB(collection)
					continue
				}
				time.Sleep(7 * time.Second)
				updateDevicesToMongoDB(collection)
				log.Info("Updated devices to MongoDB")
			}
		}

	}()
}

// Function to update changes to MongoDB
func updateDevicesToMongoDB(collection *mongo.Collection) {
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
	log.Info("Updated devices to MongoDB")
}

func eraseAllMongoData(collection *mongo.Collection) {
	ctx := context.Background()

	_, err := collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		log.Println("Error deleting all devices from MongoDB:", err)
	}
}
