package service

import (
	"time"

	"github.com/Crunchy89/go-mysql/app/user/payload"
	"github.com/Crunchy89/go-mysql/app/user/repository"
	"github.com/Crunchy89/go-mysql/app/user/response"
	"github.com/Crunchy89/go-mysql/domain"
	"github.com/Crunchy89/go-mysql/utils/r"
)

type UserService interface {
	GetByUUID(uuid string) (*payload.UserResponse, r.Ex)
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

func (b *baseUserService) GetByUUID(uuid string) (*payload.UserResponse, r.Ex) {
	res, err := b.user.GetByUUID(uuid)
	response := response.SingleResponse(res)
	return response, err
}
func (b *baseUserService) GetAll() ([]*payload.UserResponse, r.Ex) {
	res, err := b.user.GetAll()
	response := response.BatchResponse(res)
	return response, err
}
func (b *baseUserService) CreateUser(user *payload.UserCreate) (string, r.Ex) {
	newUser := &domain.User{
		Username: user.Username,
		Password: user.Password,
		RoleID:   user.RoleID,
	}
	_, err := b.user.CreateUser(newUser)
	if err != nil {
		return "", err
	}
	return "data created", nil
}
func (b *baseUserService) UpdateUser(user *payload.UserUpdate) (string, r.Ex) {
	current, err := b.user.GetByUUID(user.UUID)
	if err != nil {
		return "", err
	}
	if user.Password != nil {
		current.Password = *user.Password
	}
	if user.RoleID != nil {
		current.RoleID = *user.RoleID
	}
	current.UpdatedAt = time.Now().UTC()
	return b.user.UpdateUser(current)
}
func (b *baseUserService) DeleteUser(user *payload.UserDelete) (string, r.Ex) {
	current, err := b.user.GetByUUID(user.UUID)
	if err != nil {
		return "", err
	}
	return b.user.DeleteUser(current)
}
