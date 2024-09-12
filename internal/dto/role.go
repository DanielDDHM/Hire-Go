package models

type Role struct {
	Id          int    `gorm:"type:int;not null" json:"id"`
	Name        string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	Description string `gorm:"type:varchar(100);not null" json:"description" validate:"required"`
}
