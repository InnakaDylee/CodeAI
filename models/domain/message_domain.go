package domain

import "time"

type Message struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        int64 `gorm:"primarykey" gorm:"autoIncrement"`
	UserID    int64
	Name      string
	Text      string
}