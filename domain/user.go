package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:50;unique"`
	Password string `gorm:"size:255"`
	UUID     string `gorm:"size:255"`
	// define foreign key on update cascade on delete cascade
	RoleID uint
	Role   Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
