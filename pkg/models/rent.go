package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	Id     uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CartId uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	BookId uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}
