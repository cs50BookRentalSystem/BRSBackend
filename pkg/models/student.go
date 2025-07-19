package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model `json:"-"`
	Id         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:(gen_random_uuid())"`
	FirstName  string    `json:"first_name" validate:"required" gorm:"type:varchar(255);not null"`
	LastName   string    `json:"last_name" validate:"required" gorm:"type:varchar(255);not null"`
	CardId     string    `json:"card_id" gorm:"type:varchar(255);not null"`
	Major      string    `json:"major" validate:"required" gorm:"type:varchar(255);not null"`
	Phone      string    `json:"phone" validate:"required" gorm:"type:varchar(255);not null"`
}
