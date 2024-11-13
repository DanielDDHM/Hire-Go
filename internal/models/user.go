package models

import "time"

type User struct {
	Id           int       `gorm:"type:int;not null" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	ProfilePhoto string    `gorm:"type:varchar(100);not null" json:"profile_photo" validate:"omitempty,url"`
	Email        string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email" validate:"required,email"`
	Password     string    `gorm:"type:varchar(255);not null" json:"-" validate:"required"`
	RoleId       int       `gorm:"type:int;not null" json:"role_id" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
