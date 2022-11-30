package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Role string `gorm:"size:50;unique"`
	UUID string `gorm:"size:255"`
}

func (u *Role) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New().String()
	return
}
