package handler

import (
	"code-ai/services"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
)

type BotHandler struct {
	BotService services.BotService
    OpenAIService services.CodeService
}

func NewBotHandler(botService services.BotService, openAiService services.CodeService) *BotHandler {
	return &BotHandler{
		BotService: botService,
        OpenAIService: openAiService,
	}
}

func (wh *BotHandler) HandleBot(c echo.Context)error{
	update := new(tgbotapi.Update)
    if err := c.Bind(update); err != nil {
        return err
    }

    // Handle the update here using Bot Service
    if update.Message != nil {
        answer, _ := wh.OpenAIService.CodeGenerator(update.Message.Text)
        msg, err := wh.BotService.SendMessage(update.Message.Chat.ID, answer)
        if err != nil {
            return err
        }
        // Handle the response message returned from the service
        return c.JSON(http.StatusOK, msg)
    }

    return c.String(http.StatusOK, "OK")
}