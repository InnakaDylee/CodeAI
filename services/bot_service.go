package services

import (
	"code-ai/models/domain"
	"code-ai/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotService struct {
    Repository repository.BotRepository
    Bot        *tgbotapi.BotAPI
}

type BotServiceInterface interface {
    SendMessage(chatID int64, text string) (*domain.Message, error)
    InitializeBot(token string) error
	SaveMessage(chatID int64, name, text string) (*domain.Message, error)	
}

func NewBotService(repo repository.BotRepository, token string) *BotService {
    bot, _ := tgbotapi.NewBotAPI(token)
    return &BotService{
		Repository: repo,
		Bot: bot,
	}
}

func (ws *BotService) SendMessage(chatID int64, text string) (*domain.Message, error) {
	msg := tgbotapi.NewMessage(chatID, text)
    _, err := ws.Bot.Send(msg)
    if err != nil {
        return nil, err
    }

	ws.Repository.SaveMessage(ws.Bot.Self.ID, ws.Bot.Self.FirstName, text)

    return &domain.Message{ID: chatID, Text: text}, nil
}

func (ws *BotService) InitializeBot(token string) error {
    bot, err := tgbotapi.NewBotAPI(token)
    if err != nil {
        return err
    }
    ws.Bot = bot
    return nil
}

func (ws *BotService) SaveMessage(chatID int64, name, text string) (*domain.Message, error) {
	message, err := ws.Repository.SaveMessage(chatID, name, text)
	if err != nil {
		return nil, err
	}
	return message, nil
}