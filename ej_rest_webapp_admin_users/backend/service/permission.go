package service

import (
	"backend/domain"
	"backend/dto"
	"backend/errs"
)

type PermissionService interface {
	Permissions() (*[]dto.ResponsePermission, *errs.AppError)
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
