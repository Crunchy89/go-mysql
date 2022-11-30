package domain

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Role string `gorm:"size:50;unique"`
	UUID string `gorm:"size:255"`
}
