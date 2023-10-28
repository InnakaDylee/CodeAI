package domain

import "time"

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        int64
	Name      string 
	Status    string
	TextCount int64
}