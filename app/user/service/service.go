package service

import (
	"time"

	"github.com/Crunchy89/go-mysql/app/user/payload"
	"github.com/Crunchy89/go-mysql/app/user/repository"
	"github.com/Crunchy89/go-mysql/app/user/response"
	"github.com/Crunchy89/go-mysql/domain"
	"github.com/Crunchy89/go-mysql/utils/r"
	"github.com/google/uuid"
)

type UserService interface {
	GetById(id int) (*payload.UserResponse, r.Ex)
	GetAll() ([]*payload.UserResponse, r.Ex)
	CreateUser(user *payload.UserCreate) (string, r.Ex)
	UpdateUser(user *payload.UserUpdate) (string, r.Ex)
	DeleteUser(user *payload.UserDelete) (string, r.Ex)
}

type baseUserService struct {
	user repository.UserRepository
}

func NewUserService(user repository.UserRepository) UserService {
	return &baseUserService{user: user}
}

func (b *baseUserService) GetById(id int) (*payload.UserResponse, r.Ex) {
	res, err := b.user.GetByID(id)
	response := response.SingleResponse(res)
	return response, err
}
func (b *baseUserService) GetAll() ([]*payload.UserResponse, r.Ex) {
	res, err := b.user.GetAll()
	response := response.BatchResponse(res)
	return response, err
}
func (b *baseUserService) CreateUser(user *payload.UserCreate) (string, r.Ex) {
	UUID := uuid.New().String()
	newUser := &domain.User{
		Username: user.Username,
		Password: user.Password,
		UUID:     UUID,
		RoleID:   user.RoleID,
	}
	_, err := b.user.CreateUser(newUser)
	if err != nil {
		return "", err
	}
	return "data created", nil
}
func (b *baseUserService) UpdateUser(user *payload.UserUpdate) (string, r.Ex) {
	newUser := &domain.User{}
	newUser.ID = user.ID
	if user.Password != nil {
		newUser.Password = *user.Password
	}
	if user.RoleID != nil {
		newUser.RoleID = *user.RoleID
	}
	newUser.UpdatedAt = time.Now().UTC()
	return b.user.UpdateUser(newUser)
}
func (b *baseUserService) DeleteUser(user *payload.UserDelete) (string, r.Ex) {
	newUser := &domain.User{}
	newUser.ID = user.ID
	return b.user.DeleteUser(newUser)
}
