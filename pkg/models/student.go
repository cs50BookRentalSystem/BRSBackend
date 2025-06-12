package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model `json:"-"`
	Id         uuid.UUID `gorm:"primaryKey;type:uuid;default:(gen_random_uuid())" json:"id"`
	CardId     string    `gorm:"type:varchar(255);not null" json:"card_id"`
	FirstName  string    `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName   string    `gorm:"type:varchar(255);not null" json:"last_name"`
	Major      string    `gorm:"type:varchar(255);not null" json:"major"`
	Phone      string    `gorm:"type:varchar(255);not null" json:"phone"`
}

//func (b *Student) BeforeCreate(tx *gorm.DB) (err error) {
//	b.Id, err = uuid.NewUUID()
//	if err != nil {
//		return err
//	}
//	return nil
//}
