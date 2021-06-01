package domain

import (
	"backend/dto"
	"backend/errs"
)

type User struct {
	ID        uint   `gorm:"column:id;primaryKey"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email;unique"`
	Password  string `gorm:"column:password"`
	RoleID    uint   `gorm:"column:role_id"`
	Role      Role   `gorm:"foreignKey:role_id"`
}

type UserStorage interface {
	InsertUser(User) (*User, *errs.AppError)
	SelectByLogin(User) (*User, *errs.AppError)
	SelectUser(User) (*User, *errs.AppError)
	SelectUsers(int) (*[]User, int64, *errs.AppError)
	UpdateUser(User) (*User, *errs.AppError)
	DeleteUser(User) *errs.AppError
}

func NewUser(id uint, fn, ln, em, pass string, rol uint) User {
	return User{
		ID:        id,
		FirstName: fn,
		LastName:  ln,
		Email:     em,
		Password:  pass,
		RoleID:    rol,
	}
}

func (u User) ToDto() dto.ResponseUser {
	return dto.ResponseUser{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		RoleID:    u.RoleID,
		Role: dto.ResponseRole{
			ID:   u.Role.ID,
			Name: u.Role.Name,
		},
	}
}
