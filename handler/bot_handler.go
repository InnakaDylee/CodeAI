package handler

import (
	"code-ai/models/domain"
	"code-ai/services"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
)

type BotHandler struct {
	BotService services.BotService
    OpenAIService services.CodeService
	User services.UserServiceImp
}

func NewBotHandler(botService services.BotService, openAiService services.CodeService, user services.UserServiceImp) *BotHandler {
	return &BotHandler{
		BotService: botService,
        OpenAIService: openAiService,
		User: user,
	}
}

func (wh *BotHandler) HandleBot(c echo.Context)error{
	update := new(tgbotapi.Update)
    if err := c.Bind(update); err != nil {
        return err
    }

    // Handle the update here using Bot Service
    if update.Message != nil {
		user, err := wh.User.FindUserByID(update.Message.From.ID)
		if err != nil {
			createUser := domain.User{
				ID: update.Message.From.ID,
				Name: update.Message.From.UserName,
				Status: "UNPAID",
				TextCount: 0,
			}
			user, _ = wh.User.CreateUser(&createUser)
		}
		if user.Status == "UNPAID" && user.TextCount > 100 {
			wh.BotService.SendMessage(update.Message.Chat.ID, "Mohon maaf, anda sudah melebihi batas penggunaan. Silahkan upgrade ke premium untuk menggunakan layanan ini.")
			return c.String(http.StatusOK, "OK")
		}
		wh.User.AddLimitText(update.Message.From.ID, user.TextCount)
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