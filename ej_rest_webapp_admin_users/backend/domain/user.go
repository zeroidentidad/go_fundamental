package domain

import "backend/errs"

type User struct {
	ID        uint   `gorm:"primaryKey,column:id"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
}

type UserStorage interface {
	InsertUser(User) (*User, *errs.AppError)
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
