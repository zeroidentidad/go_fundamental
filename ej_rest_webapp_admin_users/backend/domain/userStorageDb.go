package domain

import (
	"backend/errs"
	"backend/logs"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserStorageDb struct {
	client *gorm.DB
}

func NewUserStorageDb(dbClient *gorm.DB) UserStorageDb {
	return UserStorageDb{dbClient}
}

func (db UserStorageDb) InsertUser(u User) (*User, *errs.AppError) {
	r := db.client.Create(&u)
	if r.Error != nil {
		logs.Error("Error creating user: " + r.Error.Error())
		return &u, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &u, nil
}

func (db UserStorageDb) SelectByLogin(u User) (*User, *errs.AppError) {
	var usr User
	r := db.client.Where("email = ?", u.Email).First(&usr)
	if r.Error != nil {
		if errors.Is(r.Error, gorm.ErrRecordNotFound) {
			return &usr, errs.NewUnexpectedError("User not found")
		}
		logs.Error("Error finding user: " + r.Error.Error())
		return &usr, errs.NewUnexpectedError("Unexpected error from database")
	}

	passwordReq := []byte(u.Password)
	passwordDB := []byte(usr.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordReq)
	if err != nil {
		return &usr, errs.NewBadRequestError("Wrong credentials")
	}

	return &usr, nil
}

func (db UserStorageDb) SelectUser(u User) (*User, *errs.AppError) {
	var usr User
	r := db.client.Where("id = ?", u.ID).First(&usr)
	if r.Error != nil {
		if errors.Is(r.Error, gorm.ErrRecordNotFound) {
			return &usr, errs.NewUnexpectedError("User not found")
		}
		logs.Error("Error finding user: " + r.Error.Error())
		return &usr, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &usr, nil
}
