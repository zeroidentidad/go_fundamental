package domain

import (
	"backend/dto"
	"backend/errs"
)

type Role struct {
	ID   uint   `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}

type RoleStorage interface {
	InsertRole(Role) (*Role, *errs.AppError)
	SelectRole(Role) (*Role, *errs.AppError)
	SelectRoles() (*[]Role, *errs.AppError)
	UpdateRole(Role) (*Role, *errs.AppError)
	DeleteRole(Role) *errs.AppError
}

func NewRole(id uint, name string) Role {
	return Role{
		ID:   id,
		Name: name,
	}
}

func (r Role) ToDto() dto.ResponseRole {
	return dto.ResponseRole{
		ID:   r.ID,
		Name: r.Name,
	}
}
