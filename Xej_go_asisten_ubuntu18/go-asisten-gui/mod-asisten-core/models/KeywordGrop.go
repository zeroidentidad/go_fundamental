package models

import (
	"github.com/jinzhu/gorm"
)

// KeywordGroup models
type KeywordGroup struct {
	gorm.Model
	KeyWordID  uint
	KeyWord    *KeyWord `gorm:"foreignkey:ID;association_foreignkey:KeyWordID"`
	GroupKeyID uint
}
