package models

import "github.com/jinzhu/gorm"

// Action is response
type Action struct {
	gorm.Model
	Name     string
	SystemGo string
	Command  string
}
