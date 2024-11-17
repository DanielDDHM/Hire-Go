package models

import "time"

type Model struct {
	Id        int       `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	UserId    int       `gorm:"type:int;not null" json:"user_id" validate:"required"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name" validate:"required"`
	Age       *int      `gorm:"type:int" json:"age" validate:"omitempty,gte=0"`
	Height    *float64  `gorm:"type:decimal(4,2)" json:"height" validate:"omitempty"`
	Weight    *float64  `gorm:"type:decimal(4,2)" json:"weight" validate:"omitempty"`
	Bio       string    `gorm:"type:text" json:"bio" validate:"omitempty"`
	Photos    []string  `gorm:"type:text[]" json:"photos" validate:"omitempty,dive,url"`
	Address   string    `gorm:"type:varchar(50)" json:"address" validate:"omitempty"`
	City      string    `gorm:"type:varchar(25)" json:"city" validate:"omitempty"`
	Country   string    `gorm:"type:varchar(25)" json:"country" validate:"omitempty"`
	CreatedAt time.Time `gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
