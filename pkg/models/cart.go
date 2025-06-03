package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	StudentId uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Status    string    `gorm:"type:text;default:'RENTED'"`
}
