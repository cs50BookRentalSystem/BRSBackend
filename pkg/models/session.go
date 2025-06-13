package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model  `json:"-"`
	Id          string    `gorm:"primaryKey;type:varchar(255)" json:"id"`
	LibrarianId uuid.UUID `gorm:"type:uuid;not null;index" json:"librarian_id"`
	ExpiresAt   time.Time `gorm:"not null" json:"expires_at"`
	Librarian   Librarian `gorm:"foreignKey:LibrarianId;references:Id" json:"-"`
}
