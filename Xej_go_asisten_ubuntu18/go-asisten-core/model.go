package main

import (
	"./readconfig"
	"./systemgo"
	"github.com/jinzhu/gorm"

	"./texttospeech"
)

// AsistenCoreModel model
type AsistenCoreModel struct {
	modelAsisten             *readconfig.ModelAsisten
	speechConfg              *readconfig.ModelSpeechGroup
	clientts                 *texttospeech.Client
	system                   *systemgo.SystemGoModel
	modelCommandKeys         *readconfig.ModelCommandKeys
	modelProgramsAlternative *readconfig.ModelProgramAlternative
	Db                       *gorm.DB
	cfgDir                   string
}

// Parse func
func (ctx *AsistenCoreModel) Parse(speech string) func(string) string {
	return func(words string) string {
		words = ctx.modelProgramsAlternative.Parse(words)
		words = ctx.modelAsisten.Parse(words, speech)
		words = ctx.modelCommandKeys.Parse(words)
		return words
	}
}
