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
}

type UserStorage interface {
	InsertUser(User) (*User, *errs.AppError)
	SelectByLogin(User) (*User, *errs.AppError)
	SelectUser(u User) (*User, *errs.AppError)
	SelectUsers() (*[]User, *errs.AppError)
}

func NewUser(id uint, fn, ln, em, pass string) User {
	return User{
		ID:        id,
		FirstName: fn,
		LastName:  ln,
		Email:     em,
		Password:  pass,
	}
}

func (u User) ToDto() dto.ResponseUser {
	return dto.ResponseUser{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
}
