package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model  `json:"-"`
	Id          uuid.UUID `gorm:"primaryKey;type:uuid;default:(gen_random_uuid())" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Count       int       `gorm:"type:int;not null" json:"count"`
}
