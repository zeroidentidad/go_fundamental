package domain

import (
	"backend/errs"
	"backend/logs"
	"errors"

	"gorm.io/gorm"
)

type PermissionStorageDb struct {
	client *gorm.DB
}

func NewPermissionStorageDb(dbClient *gorm.DB) PermissionStorageDb {
	return PermissionStorageDb{dbClient}
}

func (db PermissionStorageDb) SelectPermissions() (*[]Permission, *errs.AppError) {
	permissions := make([]Permission, 0)
	r := db.client.Find(&permissions)
	if r.Error != nil {
		if errors.Is(r.Error, gorm.ErrEmptySlice) {
			return &permissions, errs.NewUnexpectedError("Permissions not found")
		}
		logs.Error("Error finding permissions: " + r.Error.Error())
		return &permissions, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &permissions, nil
}
