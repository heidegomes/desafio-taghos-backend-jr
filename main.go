package main

import (
	"desafio-taghos-backend-jr/config"
	"desafio-taghos-backend-jr/routes"
	"log"
)

func main() {
	config.LoadEnvs()

	r := routes.SetupRouter()
	log.Println("Servidor rodando na porta 8080")
	r.Run(":8080")
}