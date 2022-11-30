package domain

import (
	"github.com/Crunchy89/go-mysql/utils/password"
	"github.com/google/uuid"
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

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New().String()
	u.Password = password.HashPassword(u.Password)
	return
}
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password = password.HashPassword(u.Password)
	return
}
