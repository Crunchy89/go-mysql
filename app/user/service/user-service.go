package service

import (
	"github.com/Crunchy89/go-mysql/app/user/repository"
	"github.com/Crunchy89/go-mysql/domain"
	"github.com/Crunchy89/go-mysql/utils/r"
)

type UserService interface {
	GetById(id int) (*domain.User, r.Ex)
	GetAll() ([]*domain.User, r.Ex)
}

type baseUserService struct {
	user repository.UserRepository
}

func NewUserService(user repository.UserRepository) UserService {
	return &baseUserService{user: user}
}

func (b *baseUserService) GetById(id int) (*domain.User, r.Ex) {
	return b.user.GetByID(id)
}
func (b *baseUserService) GetAll() ([]*domain.User, r.Ex) {
	return b.user.GetAll()
}
