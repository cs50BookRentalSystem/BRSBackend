package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Librarian struct {
	gorm.Model `json:"-"`
	Id         uuid.UUID `gorm:"primaryKey;type:uuid;default:(gen_random_uuid())" json:"id"`
	User       string    `gorm:"<-:create;uniqueIndex;type:varchar(255);not null" json:"user"`
	Pass       []byte    `gorm:"type:text;not null" json:"-"`
}

//func (b *Librarian) BeforeCreate(tx *gorm.DB) (err error) {
//	b.Id, err = uuid.NewUUID()
//	if err != nil {
//		return err
//	}
//	return nil
//}
