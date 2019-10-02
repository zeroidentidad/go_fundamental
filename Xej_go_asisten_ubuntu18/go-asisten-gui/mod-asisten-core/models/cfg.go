package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

// GetAllGrupKeys proloads
func GetAllGrupKeys(db *gorm.DB) []GroupKey {
	// Get Alls GroupKeys
	var groupKeys []GroupKey
	err := db.Find(&groupKeys).Error
	if err != nil {
		log.Println("GetAllGrupKeys1: ", err)
	}

	// find KeywordGropus
	for ind := range groupKeys {
		var kg []KeywordGroup
		db.Preload("KeyWord").Find(&kg, "group_key_id = ?", groupKeys[ind].ID)
		groupKeys[ind].KeywordGroups = kg
	}

	// Return data
	return groupKeys
}
