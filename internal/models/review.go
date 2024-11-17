package models

import "time"

type Review struct {
	Id         int       `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	ReviewerId *int      `gorm:"type:int" json:"reviewer_id" validate:"omitempty"`
	ReviewedId *int      `gorm:"type:int" json:"reviewed_id" validate:"omitempty"`
	Rating     float64   `gorm:"type:decimal(2,1);not null;check:rating >= 1 AND rating <= 5" json:"rating" validate:"required,gte=1,lte=5"`
	Comment    string    `gorm:"type:text" json:"comment" validate:"omitempty"`
	CreatedAt  time.Time `gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
