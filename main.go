package main

import (
	"code-ai/drivers"
	"code-ai/routes"
	"log"

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

	e.Logger.Fatal(e.Start(":3334"))
}