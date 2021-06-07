package service

import (
	"backend/domain"
	"backend/dto"
	"backend/errs"
	"strconv"
)

type RoleService interface {
	Roles() (*[]dto.ResponseRole, *errs.AppError)
	CreateRole(dto.RequestRole) (*dto.ResponseRole, *errs.AppError)
	GetRole(string) (*dto.ResponseRole, *errs.AppError)
	UpdateRole(dto.RequestRole) (*dto.ResponseRole, *errs.AppError)
	DeleteRole(string) *errs.AppError
}

type DefaultRoleService struct {
	repo domain.RoleStorage
}

func NewRoleService(repo domain.RoleStorage) DefaultRoleService {
	return DefaultRoleService{
		repo,
	}
}

func (s DefaultRoleService) Roles() (*[]dto.ResponseRole, *errs.AppError) {
	res := make([]dto.ResponseRole, 0)
	roles, err := s.repo.SelectRoles()
	if err != nil {
		return &res, err
	}

	for _, u := range *roles {
		res = append(res, u.ToDto())
	}

	return &res, err
}

func (s DefaultRoleService) CreateRole(req dto.RequestRole) (res *dto.ResponseRole, err *errs.AppError) {
	r := domain.NewRole(req.ID, req.Name, req.Permissions)

	role, err := s.repo.InsertRole(r)
	if err != nil {
		return res, err
	}

	dto := role.ToDto()

	return &dto, nil
}

func (s DefaultRoleService) GetRole(id string) (res *dto.ResponseRole, err *errs.AppError) {
	_id, _ := strconv.Atoi(id)
	r := domain.NewRole(uint(_id), "", []string{})

	rol, err := s.repo.SelectRole(r)
	if err != nil {
		return res, err
	}

	dto := rol.ToDto()

	return &dto, nil
}

func (s DefaultRoleService) UpdateRole(req dto.RequestRole) (res *dto.ResponseRole, err *errs.AppError) {
	r := domain.NewRole(req.ID, req.Name, req.Permissions)

	role, err := s.repo.UpdateRole(r)
	if err != nil {
		return res, err
	}

	dto := role.ToDto()

	return &dto, nil
}

func (s DefaultRoleService) DeleteRole(id string) (err *errs.AppError) {
	_id, _ := strconv.Atoi(id)
	r := domain.NewRole(uint(_id), "", []string{})

	err = s.repo.DeleteRole(r)
	if err != nil {
		return err
	}

	return nil
}
