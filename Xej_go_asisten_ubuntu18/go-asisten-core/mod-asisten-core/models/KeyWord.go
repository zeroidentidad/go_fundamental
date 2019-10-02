package models

import (
	"github.com/jinzhu/gorm"
)

// KeyWord for responses
type KeyWord struct {
	gorm.Model
	Key string `gorm:"unique"`
}
