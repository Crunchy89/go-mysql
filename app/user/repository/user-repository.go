package repository

import (
	"github.com/Crunchy89/go-mysql/domain"
	"github.com/Crunchy89/go-mysql/utils/r"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]*domain.User, r.Ex)
	GetByID(id int) (*domain.User, r.Ex)
}

type baseUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &baseUserRepository{db: db}
}

func (b *baseUserRepository) GetAll() ([]*domain.User, r.Ex) {
	user := []*domain.User{}
	result := b.db.Table("tbl_customer").First(&user)
	if result.RowsAffected > 0 {
		return user, nil
	} else {
		return nil, r.NewErr(result.Error.Error())
	}
}
func (b *baseUserRepository) GetByID(id int) (*domain.User, r.Ex) {
	user := &domain.User{}
	result := b.db.Table("tbl_customer").First(&user, id)
	if result.RowsAffected > 0 {
		return user, nil
	} else {
		return nil, r.NewErr(result.Error.Error())
	}
}
