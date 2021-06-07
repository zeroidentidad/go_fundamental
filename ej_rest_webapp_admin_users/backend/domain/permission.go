package domain

import (
	"backend/dto"
	"backend/errs"
)

type Permission struct {
	ID   uint   `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}

type PermissionStorage interface {
	SelectPermissions() (*[]Permission, *errs.AppError)
	SelectRole(Role) (*Role, *errs.AppError)
}

func (r Permission) ToDto() dto.ResponsePermission {
	return dto.ResponsePermission{
		ID:   r.ID,
		Name: r.Name,
	}
}
