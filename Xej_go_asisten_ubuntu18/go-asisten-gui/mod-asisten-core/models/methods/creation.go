package methods

import (
	"../../models"
	"github.com/jinzhu/gorm"
)

// CreateBasic element
func CreateBasic(db *gorm.DB, keywords []string, action *models.Action, responses []string) error {
	// Creo las KeyWords
	var idsKeywords []uint
	for _, val := range keywords {
		var key = new(models.KeyWord)
		err := db.Find(key, "key = ?", val).Error
		if err != nil {
			key.Key = val
			db.Create(key)

		}
		idsKeywords = append(idsKeywords, key.ID)
	}
	// Creo las Actions
	db.Create(action)

	// Creo las Response
	var idsResponses []uint
	for _, val := range responses {
		var actiResponse = new(models.ActionResponse)
		actiResponse.Response = val
		actiResponse.ActionID = action.ID
		db.Create(actiResponse)
		idsResponses = append(idsResponses, actiResponse.ID)
	}
	// Creo un groupKey
	var groupKey = new(models.GroupKey)
	groupKey.ActionID = action.ID
	db.Create(groupKey)

	// Agrego las keywords al groupkey
	for _, val := range idsKeywords {
		var keyGroup = new(models.KeywordGroup)
		keyGroup.KeyWordID = val
		keyGroup.GroupKeyID = groupKey.ID
		db.Create(keyGroup)
	}
	return nil
}
