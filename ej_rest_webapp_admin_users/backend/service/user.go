package service

import (
	"backend/domain"
	"backend/dto"
	"backend/errs"

	"github.com/dgrijalva/jwt-go"
)

type UserService interface {
	Register(dto.RequestUser) (*dto.ResponseUser, *errs.AppError)
	Login(dto.RequestUser) (*dto.ResponseUserLogin, *errs.AppError)
	AuthUser(string) (*dto.UserClaims, *errs.AppError)
}

type DefaultUserService struct {
	repo domain.UserStorage
}

func NewUserService(repo domain.UserStorage) DefaultUserService {
	return DefaultUserService{
		repo,
	}
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

func (s DefaultUserService) Login(req dto.RequestUser) (res *dto.ResponseUserLogin, err *errs.AppError) {
	u := domain.NewUser(0, "", "", req.Email, req.Password)

	usr, err := s.repo.SelectByLogin(u)
	if err != nil {
		return res, err
	}

	login := &dto.ResponseUser{
		ID:        usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	}

	tk, err := login.CreateToken()
	if err != nil {
		return res, err
	}

	res = tk
	return res, nil
}

func (s DefaultUserService) AuthUser(tk string) (*dto.UserClaims, *errs.AppError) {
	var claims dto.UserClaims
	token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(dto.TOKEN_SECRET), nil
	})
	if err != nil {
		return &claims, errs.NewValidationError("Invalid token format")
	}

	if !token.Valid {
		return &claims, errs.NewValidationError("Invalid token")
	}

	mapClaims := token.Claims.(jwt.MapClaims)
	usrClaims, err := dto.FromJwtMapClaims(mapClaims)
	if err != nil {
		return &claims, errs.NewValidationError("Error token claims")
	}
	claims = *usrClaims

	return &claims, nil
}
