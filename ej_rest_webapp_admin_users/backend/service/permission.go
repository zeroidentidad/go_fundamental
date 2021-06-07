package service

import (
	"backend/domain"
	"backend/dto"
	"backend/errs"
)

type PermissionService interface {
	Permissions() (*[]dto.ResponsePermission, *errs.AppError)
	RolePermissions(uint) (*dto.ResponseRole, *errs.AppError)
}

type DefaultPermissionService struct {
	repo domain.PermissionStorage
}

func NewPermissionService(repo domain.PermissionStorage) DefaultPermissionService {
	return DefaultPermissionService{
		repo,
	}
}

func (s DefaultPermissionService) Permissions() (*[]dto.ResponsePermission, *errs.AppError) {
	res := make([]dto.ResponsePermission, 0)
	permissions, err := s.repo.SelectPermissions()
	if err != nil {
		return &res, err
	}

	for _, p := range *permissions {
		res = append(res, p.ToDto())
	}

	return &res, err
}

func (s DefaultPermissionService) RolePermissions(id uint) (res *dto.ResponseRole, err *errs.AppError) {
	r := domain.NewRole(id, "", []string{})

	rol, err := s.repo.SelectRole(r)
	if err != nil {
		return res, err
	}

	dto := rol.ToDto()

	return &dto, nil
}
