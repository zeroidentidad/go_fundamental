package domain

import (
	"backend/errs"

	"gorm.io/gorm"
)

type UserStorageDb struct {
	client *gorm.DB
}

func NewUserStorageDb(dbClient *gorm.DB) UserStorageDb {
	return UserStorageDb{dbClient}
}

func (db UserStorageDb) InsertUser(u User) (*User, *errs.AppError) {
	// _sql := `INSERT INTO users ()`

	//_, err := db.client
	/*if err != nil {
		logs.Error("Error creating new user: " + err.Error())
		return &u, errors.UnexpectedError("Unexpected error from database")
	}*/

	return &u, nil
}
