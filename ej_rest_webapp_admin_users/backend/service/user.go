package service

import (
	"backend/domain"
	"backend/dto"
	"backend/errs"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

type UserService interface {
	AuthUser(string) (*dto.UserClaims, *errs.AppError)
	Register(dto.RequestUser) (*dto.ResponseUser, *errs.AppError)
	Login(dto.RequestUser) (*dto.ResponseUserLogin, *errs.AppError)
	User(*dto.UserClaims) (*dto.ResponseUser, *errs.AppError)
	Users() (*[]dto.ResponseUser, *errs.AppError)
	CreateUser(dto.RequestUser) (*dto.ResponseUser, *errs.AppError)
	GetUser(string) (*dto.ResponseUser, *errs.AppError)
	UpdateUser(dto.RequestUser) (*dto.ResponseUser, *errs.AppError)
	DeleteUser(string) *errs.AppError
}

type DefaultUserService struct {
	repo domain.UserStorage
}

func NewUserService(repo domain.UserStorage) DefaultUserService {
	return DefaultUserService{
		repo,
	}
}

func (s DefaultUserService) AuthUser(tk string) (*dto.UserClaims, *errs.AppError) {
	var claims dto.UserClaims
	token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(dto.TOKEN_SECRET), nil
	})
	if err != nil {
		return &claims, errs.NewBadRequestError("Without session token")
	}

	if !token.Valid {
		return &claims, errs.NewBadRequestError("Invalid session token")
	}

	mapClaims := token.Claims.(jwt.MapClaims)
	usrClaims, err := dto.FromJwtMapClaims(mapClaims)
	if err != nil {
		return &claims, errs.NewValidationError("Error token claims")
	}
	claims = *usrClaims

	return &claims, nil
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

	dto := usr.ToDto()

	return &dto, nil
}

func (s DefaultUserService) Login(req dto.RequestUser) (res *dto.ResponseUserLogin, err *errs.AppError) {
	u := domain.NewUser(0, "", "", req.Email, req.Password)

	usr, err := s.repo.SelectByLogin(u)
	if err != nil {
		return res, err
	}

	login := usr.ToDto()

	tk, err := login.CreateToken()
	if err != nil {
		return res, err
	}

	res = tk
	return res, nil
}

func (s DefaultUserService) User(req *dto.UserClaims) (res *dto.ResponseUser, err *errs.AppError) {
	u := domain.NewUser(req.ID, "", "", "", "")

	usr, err := s.repo.SelectUser(u)
	if err != nil {
		return res, err
	}

	dto := usr.ToDto()

	return &dto, nil
}

func (s DefaultUserService) Users() (*[]dto.ResponseUser, *errs.AppError) {
	res := make([]dto.ResponseUser, 0)
	users, err := s.repo.SelectUsers()
	if err != nil {
		return &res, err
	}

	for _, u := range *users {
		res = append(res, u.ToDto())
	}

	return &res, err
}

func (s DefaultUserService) CreateUser(req dto.RequestUser) (res *dto.ResponseUser, err *errs.AppError) {
	pass, err := req.EncryptPassword()
	if err != nil {
		return res, err
	}

	u := domain.NewUser(req.ID, req.FirstName, req.LastName, req.Email, pass)

	usr, err := s.repo.InsertUser(u)
	if err != nil {
		return res, err
	}

	dto := usr.ToDto()

	return &dto, nil
}

func (s DefaultUserService) GetUser(id string) (res *dto.ResponseUser, err *errs.AppError) {
	_id, _ := strconv.Atoi(id)
	u := domain.NewUser(uint(_id), "", "", "", "")

	usr, err := s.repo.SelectUser(u)
	if err != nil {
		return res, err
	}

	dto := usr.ToDto()

	return &dto, nil
}

func (s DefaultUserService) UpdateUser(req dto.RequestUser) (res *dto.ResponseUser, err *errs.AppError) {
	if req.Password != "" {
		pass, err := req.EncryptPassword()
		if err != nil {
			return res, err
		}
		req.Password = pass
	}

	u := domain.NewUser(req.ID, req.FirstName, req.LastName, req.Email, req.Password)

	usr, err := s.repo.UpdateUser(u)
	if err != nil {
		return res, err
	}

	dto := usr.ToDto()

	return &dto, nil
}

func (s DefaultUserService) DeleteUser(id string) (err *errs.AppError) {
	_id, _ := strconv.Atoi(id)
	u := domain.NewUser(uint(_id), "", "", "", "")

	err = s.repo.DeleteUser(u)
	if err != nil {
		return err
	}

	return nil
}
