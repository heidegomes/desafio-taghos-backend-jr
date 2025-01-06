package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func GetMongoURI() string {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatalf("MONGODB_URI not set in environment variables")
	}
	return mongoURI
}

func GetDB() string {
	port := os.Getenv("MONGODB_DB")
	if port == "" {
		port = "8080" 
	}
	return port
}

func GetBooks() string {
	env := os.Getenv("MONGODB_COLLECTION")
	if env == "" {
		env = "collection"
	}
	return env
}
