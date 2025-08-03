package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model `json:"-"`
	Id         uuid.UUID `gorm:"primaryKey;type:uuid;default:(gen_random_uuid())" json:"id"`
	CartId     uuid.UUID `gorm:"type:uuid;" json:"cart_id"`
	BookId     uuid.UUID `gorm:"type:uuid;" json:"book_id"`
}
