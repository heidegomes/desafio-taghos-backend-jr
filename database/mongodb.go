package database

import (
	"context"
	"log"
	"time"

	"desafio-taghos-backend-jr/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// ConnectMongoDB establishes the connection to MongoDB.
func ConnectMongoDB() {
	// Obtenha o mongoURI do arquivo de configuração
	mongoURI := config.GetMongoURI()
	// Configure the MongoDB connection URI
	clientOptions := options.Client().ApplyURI(mongoURI) // Update URI if necessary

	// Establish connection with MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a new MongoDB client
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	// Test the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")
	MongoClient = client
}

// DisconnectMongoDB closes the connection to MongoDB.
func DisconnectMongoDB() {
	if MongoClient != nil {
		if err := MongoClient.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
		log.Println("Disconnected from MongoDB.")
	}	
}
