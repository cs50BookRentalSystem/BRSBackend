package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CardId    string    `gorm:"type:varchar(255);not null"`
	FirstName string    `gorm:"type:varchar(255);not null"`
	LastName  string    `gorm:"type:varchar(255);not null"`
	Major     string    `gorm:"type:varchar(255);not null"`
	Phone     string    `gorm:"type:varchar(255);not null"`
}
