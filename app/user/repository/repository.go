package repository

import (
	"github.com/Crunchy89/go-mysql/domain"
	"github.com/Crunchy89/go-mysql/utils/r"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]*domain.User, r.Ex)
	GetByUUID(uuid string) (*domain.User, r.Ex)
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
	result := b.db.Find(&user)
	if result.RowsAffected > 0 {
		return user, nil
	}
	return nil, r.NewErr(result.Error.Error())

}
func (b *baseUserRepository) GetByUUID(uuid string) (*domain.User, r.Ex) {
	user := &domain.User{}
	result := b.db.First(&user, "uuid = ?", uuid)
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
	return nil, r.NewErr(result.Error.Error())

}
func (b *baseUserRepository) UpdateUser(user *domain.User) (string, r.Ex) {
	result := b.db.Model(&user).UpdateColumns(domain.User{RoleID: user.RoleID, Password: user.Password})
	if result.RowsAffected > 0 {
		return "data update successfully", nil
	}
	return "", r.NewErr(result.Error.Error())

}
func (b *baseUserRepository) DeleteUser(user *domain.User) (string, r.Ex) {
	result := b.db.Delete(user)
	if result.RowsAffected > 0 {
		return "data deleted", nil
	}
	return "", r.NewErr(result.Error.Error())

}
