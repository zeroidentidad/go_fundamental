package domain

import (
	"backend/errs"
	"backend/logs"

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
