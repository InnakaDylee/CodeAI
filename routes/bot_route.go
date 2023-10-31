package routes

import (
	"code-ai/handler"
	"code-ai/repository"
	"code-ai/services"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BotRouteInit(e *echo.Echo, DB *gorm.DB){
	botToken := os.Getenv("BOT_TOKEN")
    if botToken == "" {
        log.Fatal("Token bot tidak ditemukan. Pastikan Anda mengatur BOT_TOKEN.")
    }
	openAiToken := os.Getenv("OPENAI_TOKEN")
	if openAiToken == "" {
        log.Fatal("Token open ai tidak ditemukan. Pastikan Anda mengatur OPENAI_TOKEN.")
    }

	botRepo := repository.NewBotRepository(DB)
	userRepo := repository.NewUserRepository(DB)
    botService := services.NewBotService(*botRepo ,botToken)
	openAiService := services.NewCodeService(openAiToken)
	userService := services.NewUserService(*userRepo)
    handler := handler.NewBotHandler(*botService,openAiService,userService)

	e.GET("/",func(c echo.Context) error {
		return c.JSON(200,map[string]interface{}{
			"status":"ok",
			"code":200,
		})
	})
	e.POST("/webhook",handler.HandleBot)

}