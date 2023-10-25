package routes

import (
	"code-ai/handler"
	"code-ai/repository"
	"code-ai/services"
	"code-ai/drivers"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func RouteInit(){
	botToken := os.Getenv("BOT_TOKEN")
    if botToken == "" {
        log.Fatal("Token bot tidak ditemukan. Pastikan Anda mengatur BOT_TOKEN.")
    }
	openAiToken := os.Getenv("OPENAI_TOKEN")
	if openAiToken == "" {
        log.Fatal("Token open ai tidak ditemukan. Pastikan Anda mengatur OPENAI_TOKEN.")
    }

	repo := repository.NewBotRepository(drivers.DB)
    botService := services.NewBotService(*repo ,botToken)
	openAiService := services.NewCodeService(openAiToken)
    handler := handler.NewBotHandler(botService,openAiService)

	e := echo.New()

	e.GET("/",func(c echo.Context) error {
		return c.JSON(200,map[string]interface{}{
			"status":"ok",
			"code":200,
		})
	})
	e.POST("/webhook",handler.HandleBot)

	e.Logger.Fatal(e.Start(":3333"))
}