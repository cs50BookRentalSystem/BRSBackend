package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model  `json:"-"`
	Id          uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:(gen_random_uuid())"`
	Title       string    `json:"title" validate:"required" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	Count       int       `json:"count" validate:"min=0" gorm:"type:int;not null"`
}
