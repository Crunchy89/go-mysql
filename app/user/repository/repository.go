package repository

import (
	"github.com/Crunchy89/go-mysql/domain"
	"github.com/Crunchy89/go-mysql/utils/r"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]*domain.User, r.Ex)
	GetByID(id int) (*domain.User, r.Ex)
	CreateUser(user *domain.User) (*uint, r.Ex)
	UpdateUser(user *domain.User) (string, r.Ex)
	DeleteUser(user *domain.User) (string, r.Ex)
}

type baseUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &baseUserRepository{db: db}
}

func (b *baseUserRepository) GetAll() ([]*domain.User, r.Ex) {
	user := []*domain.User{}
	result := b.db.First(&user)
	if result.RowsAffected > 0 {
		return user, nil
	}
	return nil, r.NewErr(result.Error.Error())

}
func (b *baseUserRepository) GetByID(id int) (*domain.User, r.Ex) {
	user := &domain.User{}
	result := b.db.First(&user, id)
	if result.RowsAffected > 0 {
		return user, nil
	}
	return nil, r.NewErr(result.Error.Error())

}
func (b *baseUserRepository) CreateUser(user *domain.User) (*uint, r.Ex) {
	result := b.db.Create(user)
	if result.RowsAffected > 0 {
		return &user.ID, nil
	}
	return nil, r.NewErrorDatabase(result.Error)

}
func (b *baseUserRepository) UpdateUser(user *domain.User) (string, r.Ex) {
	result := b.db.First(&user).Save(user)
	if result.RowsAffected > 0 {
		return "data update successfully", nil
	}
	return "", r.NewErrorDatabase(result.Error)

}
func (b *baseUserRepository) DeleteUser(user *domain.User) (string, r.Ex) {
	result := b.db.Delete(user)
	if result.RowsAffected > 0 {
		return "data deleted", nil
	}
	return "", r.NewErrorDatabase(result.Error)

}
