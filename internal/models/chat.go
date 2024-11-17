package models

import (
	"time"

	"gorm.io/datatypes"
)

type Chat struct {
	Id           int            `gorm:"type:int;not null" json:"-"`
	Participants datatypes.JSON `gorm:"type:jsonb" json:"participants"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type Message struct {
	Id        int       `gorm:"type:int;not null" json:"-"`
	ChatId    uint      `gorm:"index"`
	SenderId  int       `gorm:"type:int;not null" json:"sender_id" validate:"required"`
	Content   string    `gorm:"type:varchar(255);not null" json:"content" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
