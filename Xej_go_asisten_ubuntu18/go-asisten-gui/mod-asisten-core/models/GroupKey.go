package models

import (
	"strings"

	"github.com/jinzhu/gorm"
)

// GroupKey model
type GroupKey struct {
	gorm.Model
	ActionID      uint
	KeywordGroups []KeywordGroup `gorm:"foreignkey:GroupKeyID"`
}

// GroupKeyValidWords  Words methods
func (ctx *GroupKey) GroupKeyValidWords(words string) (bool, int) {
	var keywords []string
	for _, val := range ctx.KeywordGroups {
		keywords = append(keywords, val.KeyWord.Key)
	}
	return containsAll(words, keywords...), len(ctx.KeywordGroups)
}

func containsAll(words string, arr ...string) bool {
	words = strings.ToLower(words)
	var longi = len(arr)
	var encontrados = 0
	for _, val := range arr {
		if strings.Contains(words, strings.ToLower(val)) {
			encontrados++
		}
	}
	return longi == encontrados
}
