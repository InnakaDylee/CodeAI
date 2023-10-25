package repository

import (
	"github.com/jinzhu/gorm"
)

type BotRepository struct {
    db *gorm.DB
}

func NewBotRepository(db *gorm.DB) *BotRepository {
    return &BotRepository{
        db,
    }
}