package models

import (
	"github.com/jinzhu/gorm"
)

// CommandKey Models
type CommandKey struct {
	gorm.Model
	Key     string
	Command string
}
