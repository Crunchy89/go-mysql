package service

import (
	"time"

	"github.com/Crunchy89/go-mysql/app/role/payload"
	"github.com/Crunchy89/go-mysql/app/role/repository"
	"github.com/Crunchy89/go-mysql/app/role/response"
	"github.com/Crunchy89/go-mysql/domain"
	"github.com/Crunchy89/go-mysql/utils/r"
	"github.com/google/uuid"
)

type RoleService interface {
	GetById(id int) (*payload.RoleResponse, r.Ex)
	GetAll() ([]*payload.RoleResponse, r.Ex)
	CreateRole(role *payload.RoleCreate) (string, r.Ex)
	UpdateRole(role *payload.RoleUpdate) (string, r.Ex)
	DeleteRole(role *payload.RoleDelete) (string, r.Ex)
}

type baseRoleService struct {
	role repository.RoleRepository
}

func NewRoleService(role repository.RoleRepository) RoleService {
	return &baseRoleService{role: role}
}

func (b *baseRoleService) GetById(id int) (*payload.RoleResponse, r.Ex) {
	res, err := b.role.GetByID(id)
	if err != nil {
		return nil, err
	}
	response := response.SingleResponse(res)
	return response, nil

}
func (b *baseRoleService) GetAll() ([]*payload.RoleResponse, r.Ex) {
	res, err := b.role.GetAll()
	if err != nil {
		return nil, err
	}
	response := response.BatchResponse(res)
	return response, nil
}
func (b *baseRoleService) CreateRole(role *payload.RoleCreate) (string, r.Ex) {
	UUID := uuid.New().String()
	newRole := &domain.Role{
		Role: role.Role,
		UUID: UUID,
	}
	_, err := b.role.CreateRole(newRole)
	if err != nil {
		return "", err
	}
	return "data created", nil
}
func (b *baseRoleService) UpdateRole(role *payload.RoleUpdate) (string, r.Ex) {
	newRole := &domain.Role{}
	newRole.ID = role.ID
	newRole.Role = role.Role
	newRole.UpdatedAt = time.Now().UTC()
	return b.role.UpdateRole(newRole)
}
func (b *baseRoleService) DeleteRole(role *payload.RoleDelete) (string, r.Ex) {
	newRole := &domain.Role{}
	newRole.ID = role.ID
	return b.role.DeleteRole(newRole)
}
