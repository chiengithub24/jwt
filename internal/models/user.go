package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	Username  string `gorm:"uniqueIndex;not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	FullName  string
	Bio       string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
