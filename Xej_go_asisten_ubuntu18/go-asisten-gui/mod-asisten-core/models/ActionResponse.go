package models

import "github.com/jinzhu/gorm"

// ActionResponse for Asisten
type ActionResponse struct {
	gorm.Model
	ActionID uint
	Response string
}
