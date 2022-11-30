package repository

import (
	"github.com/Crunchy89/go-mysql/domain"
	"github.com/Crunchy89/go-mysql/utils/r"

	"gorm.io/gorm"
)

type RoleRepository interface {
	GetAll() ([]*domain.Role, r.Ex)
	GetByID(id int) (*domain.Role, r.Ex)
	CreateRole(role *domain.Role) (*uint, r.Ex)
	UpdateRole(role *domain.Role) (string, r.Ex)
	DeleteRole(role *domain.Role) (string, r.Ex)
}

type baseRoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &baseRoleRepository{db: db}
}

func (b *baseRoleRepository) GetAll() ([]*domain.Role, r.Ex) {
	role := []*domain.Role{}
	result := b.db.First(&role)
	if result.RowsAffected > 0 {
		return role, nil
	}
	return nil, r.NewErr(result.Error.Error())

}
func (b *baseRoleRepository) GetByID(id int) (*domain.Role, r.Ex) {
	role := &domain.Role{}
	result := b.db.First(&role, id)
	if result.RowsAffected > 0 {
		return role, nil
	}
	return nil, r.NewErr(result.Error.Error())

}
func (b *baseRoleRepository) CreateRole(role *domain.Role) (*uint, r.Ex) {
	result := b.db.Create(role)
	if result.RowsAffected > 0 {
		return &role.ID, nil
	}
	return nil, r.NewErrorRepository("roles", result.Error)

}
func (b *baseRoleRepository) UpdateRole(role *domain.Role) (string, r.Ex) {
	result := b.db.First(&role).Save(role)
	if result.RowsAffected > 0 {
		return "data update successfully", nil
	}
	return "", r.NewErrorRepository("roles", result.Error)

}
func (b *baseRoleRepository) DeleteRole(role *domain.Role) (string, r.Ex) {
	result := b.db.Delete(role)
	if result.RowsAffected > 0 {
		return "data deleted", nil
	}
	return "", r.NewErrorDatabase(result.Error)

}
