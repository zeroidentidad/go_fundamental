package domain

import (
	"backend/dto"
	"backend/errs"
	"strconv"
)

type Role struct {
	ID          uint         `gorm:"column:id;primaryKey"`
	Name        string       `gorm:"column:name"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}

type RoleStorage interface {
	InsertRole(Role) (*Role, *errs.AppError)
	SelectRole(Role) (*Role, *errs.AppError)
	SelectRoles() (*[]Role, *errs.AppError)
	UpdateRole(Role) (*Role, *errs.AppError)
	DeleteRole(Role) *errs.AppError
}

func NewRole(id uint, name string, perms []string) Role {
	permissions := make([]Permission, len(perms))
	for index, value := range perms {
		id, _ := strconv.Atoi(value)
		permissions[index] = Permission{
			ID: uint(id),
		}
	}
	return Role{
		ID:          id,
		Name:        name,
		Permissions: permissions,
	}
}

func (r Role) ToDto() dto.ResponseRole {
	permissions := make([]dto.ResponsePermission, len(r.Permissions))
	for index, value := range r.Permissions {
		permissions[index] = dto.ResponsePermission{
			ID:   value.ID,
			Name: value.Name,
		}
	}
	return dto.ResponseRole{
		ID:          r.ID,
		Name:        r.Name,
		Permissions: permissions,
	}
}
