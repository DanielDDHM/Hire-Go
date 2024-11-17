package models

import "time"

type Booker struct {
	Id        int       `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	UserId    int       `gorm:"type:int;not null" json:"user_id" validate:"required"`
	Name      string    `gorm:"type:varchar(255)" json:"name" validate:"omitempty"`
	Phone     string    `gorm:"type:varchar(20)" json:"phone" validate:"omitempty,e164"`
	Address   string    `gorm:"type:varchar(50)" json:"address" validate:"omitempty"`
	City      string    `gorm:"type:varchar(25)" json:"city" validate:"omitempty"`
	Country   string    `gorm:"type:varchar(25)" json:"country" validate:"omitempty"`
	CreatedAt time.Time `gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
