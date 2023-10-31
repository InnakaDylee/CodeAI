package repository

import (
	"code-ai/models/domain"

	"gorm.io/gorm"
)

type BotRepository struct {
	db *gorm.DB
}

type BotRepositoryInterface interface {
	SaveMessage(chatID int64, name, text string) (*domain.Message, error)
}

func NewBotRepository(db *gorm.DB) *BotRepository {
	return &BotRepository{
		db,
	}
}

func (r *BotRepository) SaveMessage(chatID int64, name, text string) (*domain.Message, error) {
	message := domain.Message{
		UserID: chatID,
		Name:   name,
		Text:   text,
	}
	err := r.db.Create(&message).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}
