package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model `json:"-"`
	Id         uuid.UUID `gorm:"primaryKey;type:uuid;default:(gen_random_uuid())" json:"id"`
	StudentId  uuid.UUID `gorm:"type:uuid;" json:"student_id"`
	Status     string    `gorm:"type:text;default:'RENTED'" json:"status"`
}

//func (b *Cart) BeforeCreate(tx *gorm.DB) (err error) {
//	b.Id, err = uuid.NewUUID()
//	if err != nil {
//		return err
//	}
//	return nil
//}
