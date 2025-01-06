package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadCarregarVariaveis carrega as variáveis de ambiente a partir do arquivo .env
func LoadEnvs() {
	// Carrega as variáveis do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// GetMongoURI retorna a URL de conexão do MongoDB
func GetMongoURI() string {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatalf("MONGODB_URI not set in environment variables")
	}
	return mongoURI
}

// GetPort retorna a porta configurada na variável de ambiente PORT
func GetDB() string {
	port := os.Getenv("MONGODB_DB")
	if port == "" {
		port = "8080" // valor padrão
	}
	return port
}

// GetEnvironment retorna o ambiente de execução (dev, prod, etc.)
func GetBooks() string {
	env := os.Getenv("MONGODB_COLLECTION")
	if env == "" {
		env = "collection" // valor padrão
	}
	return env
}
