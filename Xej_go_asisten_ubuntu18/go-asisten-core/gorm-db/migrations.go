package db

import (
	"../mod-asisten-core/models"
	"github.com/jinzhu/gorm"
)

// Migrations  function
func Migrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.KeyWord{},
		&models.GroupKey{},
		&models.Action{},
		&models.GroupKey{},
		&models.ActionResponse{},
		&models.KeywordGroup{},
		&models.ProgramsAlternative{},
		&models.CommandKey{},
	).Error
}
