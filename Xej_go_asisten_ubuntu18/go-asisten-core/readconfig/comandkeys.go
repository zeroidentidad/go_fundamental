package readconfig

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"../mod-asisten-core/models"
	"../osfunc"
)

// ModelCommandKeys func
type ModelCommandKeys struct {
	Db *gorm.DB
}

// NewCommandKeyModel func
func NewCommandKeyModel(Db *gorm.DB) *ModelCommandKeys {
	return &ModelCommandKeys{Db: Db}
}

// Parse func
func (ctx *ModelCommandKeys) Parse(words string) string {
	var CommandKeys = new([]models.CommandKey)
	ctx.Db.Find(CommandKeys)
	var remplace []string
	for _, vals := range *CommandKeys {
		key := fmt.Sprintf("{{%s}}", vals.Key)
		if strings.Contains(words, key) {
			remplace = append(remplace, key, string(osfunc.CommandOutput(vals.Command)))
		}
	}
	replace := strings.NewReplacer(remplace...)
	return replace.Replace(words)
}
