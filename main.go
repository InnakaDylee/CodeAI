package main

import (
	"code-ai/drivers"
	"code-ai/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load .env file. Err: %s", err)
	}

	routes.RouteInit()
	drivers.Init()

}