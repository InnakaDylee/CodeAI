package services

import (
    "code-ai/models/domain"
    "code-ai/repository"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotServices struct {
    Repository repository.BotRepository
    Bot        *tgbotapi.BotAPI
}

type BotService interface {
    SendMessage(chatID int64, text string) (*domain.Message, error)
    InitializeBot(token string) error
}

func NewBotService(repo repository.BotRepository, token string) *BotServices {
    bot, _ := tgbotapi.NewBotAPI(token)
    return &BotServices{
		Repository: repo,
		Bot: bot,
	}
}

func (ws *BotServices) SendMessage(chatID int64, text string) (*domain.Message, error) {
    msg := tgbotapi.NewMessage(chatID, text)
    _, err := ws.Bot.Send(msg)
    if err != nil {
        return nil, err
    }
    return &domain.Message{ID: chatID, Text: text}, nil
}

func (ws *BotServices) InitializeBot(token string) error {
    bot, err := tgbotapi.NewBotAPI(token)
    if err != nil {
        return err
    }
    ws.Bot = bot
    return nil
}
