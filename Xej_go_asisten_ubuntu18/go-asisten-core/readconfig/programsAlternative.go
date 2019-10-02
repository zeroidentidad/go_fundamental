package readconfig

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"../mod-asisten-core/models"
)

// ModelProgramAlternative model
type ModelProgramAlternative struct {
	Db *gorm.DB
}

// NewProgramsAlternativeModel new model
func NewProgramsAlternativeModel(Db *gorm.DB) *ModelProgramAlternative {
	return &ModelProgramAlternative{Db: Db}
}

// Parse programs name
func (ctx *ModelProgramAlternative) Parse(words string) string {
	var ProgramAlternative = new([]models.ProgramsAlternative)
	ctx.Db.Find(ProgramAlternative)
	var remplace []string
	for _, vals := range *ProgramAlternative {
		name, err := vals.Exits()
		if err != nil {
			continue
		}
		key := fmt.Sprintf("{{%s}}", vals.GeneralName)
		if strings.Contains(words, key) {
			remplace = append(remplace, key, name)
		}
	}
	replace := strings.NewReplacer(remplace...)
	return replace.Replace(words)
}
