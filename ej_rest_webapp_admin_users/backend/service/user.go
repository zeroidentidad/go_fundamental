package service

import (
	"backend/domain"
	"backend/dto"
	"backend/errs"
)

type UserService interface {
	Register(dto.RequestUser) (*dto.ResponseUser, *errs.AppError)
	Login(dto.RequestUser) (*dto.ResponseUser, *errs.AppError)
}

type DefaultUserService struct {
	repo domain.UserStorage
}

func (s DefaultUserService) Register(req dto.RequestUser) (res *dto.ResponseUser, err *errs.AppError) {
	err = req.ValidatePassword()
	if err != nil {
		return res, err
	}

	pass, err := req.EncryptPassword()
	if err != nil {
		return res, err
	}

	u := domain.NewUser(req.ID, req.FirstName, req.LastName, req.Email, pass)

	usr, err := s.repo.InsertUser(u)
	if err != nil {
		return res, err
	}

	res = &dto.ResponseUser{
		ID:        usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	}

	return res, nil
}

func (s DefaultUserService) Login(req dto.RequestUser) (res *dto.ResponseUser, err *errs.AppError) {
	u := domain.NewUser(0, "", "", req.Email, req.Password)

	usr, err := s.repo.SelectByLogin(u)
	if err != nil {
		return res, err
	}

	res = &dto.ResponseUser{
		ID:        usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	}

	return res, nil
}

func NewUserService(repo domain.UserStorage) DefaultUserService {
	return DefaultUserService{
		repo,
	}
}
