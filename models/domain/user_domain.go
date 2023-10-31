package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    ID        int64 `gorm:"primarykey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *gorm.DeletedAt `gorm:"index"`
    Name      string
    Status    string
    Credit    int64
    Message   []Message `gorm:"ForeignKey:UserID;references:ID"`
}