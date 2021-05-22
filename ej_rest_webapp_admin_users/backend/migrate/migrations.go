package migrate

import (
	"backend/domain"

	"gorm.io/gorm"
)

func Migrations(database *gorm.DB) {
	_ = database.AutoMigrate(&domain.User{}, &domain.Role{})
}
