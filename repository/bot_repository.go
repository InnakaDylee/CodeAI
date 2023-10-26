package repository

import (
	"gorm.io/gorm"
)

type BotRepository struct {
    db *gorm.DB
}

func NewBotRepository(db *gorm.DB) *BotRepository {
    return &BotRepository{
        db,
    }
}