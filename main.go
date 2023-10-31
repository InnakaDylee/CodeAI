package main

import (
	"code-ai/drivers"
	"code-ai/routes"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load .env file. Err: %s", err)
	}
	e := echo.New()

	DB := drivers.InitDB()

	routes.BotRouteInit(e,DB)
	routes.AdminRouteInit(e,DB)	

	port := os.Getenv("PORT")
	setPort := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(setPort))
}