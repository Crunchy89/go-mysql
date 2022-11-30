package migration

import (
	"github.com/Crunchy89/go-mysql/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	if db.Migrator().HasTable(&domain.Role{}) {
		db.Migrator().DropTable(&domain.Role{})
	}
	db.Migrator().CreateTable(&domain.Role{})
	if db.Migrator().HasTable(&domain.User{}) {
		db.Migrator().DropTable(&domain.User{})
	}
	db.Migrator().CreateTable(&domain.User{})

}
