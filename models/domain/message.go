package domain

import "time"

type Message struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        int64	`gorm:"primarykey" gorm:"autoIncrement"`
	Name      string
	Text      string
	UserID    int64 
}