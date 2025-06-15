package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model `json:"-"`
	Id         uuid.UUID `gorm:"primaryKey;type:uuid;default:(gen_random_uuid())" json:"id"`
	CardId     string    `gorm:"type:varchar(255);not null" json:"student_card_id"`
	FirstName  string    `gorm:"type:varchar(255);not null" json:"firstName"`
	LastName   string    `gorm:"type:varchar(255);not null" json:"lastName"`
	Major      string    `gorm:"type:varchar(255);not null" json:"major"`
	Phone      string    `gorm:"type:varchar(255);not null" json:"phone"`
}
