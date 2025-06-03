package models

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Librarian struct {
	gorm.Model
	Id   uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	User string    `gorm:"<-:create;uniqueIndex;type:varchar(255);not null"`
	Pass []byte    `gorm:"type:text;not null"`
}

func NewUserId() uuid.UUID {
	userId, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("fail to create uuid")
	}
	return userId
}
