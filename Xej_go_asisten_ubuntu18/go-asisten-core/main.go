package main

import (
	"log"
	"math/rand"
	"time"

	db "./gorm-db"
	"./readconfig"
	"./systemgo"
	"./texttospeech"
)

func main() {}

// NewGoAsistenCore create GoAsistenCore Clients
func NewGoAsistenCore(DataDir string) (interface{}, error) {
	rand.Seed(time.Now().Unix())
	system := systemgo.NewSystemGoModel()
	modelAsisten, err := readconfig.ReadAsisten(DataDir)
	if err != nil {
		return nil, err
	}
	Db, err := db.OpenConnectionDB(DataDir, modelAsisten.Lang)
	if err != nil {
		return nil, err
	}
	db.Migrations(Db)
	speechGroup := readconfig.NewSpeech(Db)
	clientts := texttospeech.NewClient(modelAsisten.Lang)
	modelCommandKeys := readconfig.NewCommandKeyModel(Db)
	modelProgramsAlternative := readconfig.NewProgramsAlternativeModel(Db)
	return &AsistenCoreModel{
		Db:                       Db,
		modelAsisten:             modelAsisten,
		speechConfg:              speechGroup,
		clientts:                 clientts,
		system:                   system,
		modelCommandKeys:         modelCommandKeys,
		modelProgramsAlternative: modelProgramsAlternative,
	}, nil
}

// WordsToAsistenResponse parse
func (ctx *AsistenCoreModel) WordsToAsistenResponse(words string) (string, func()) {
	if sl := ctx.speechConfg.Find(words); sl != nil {
		parse := ctx.Parse(words)
		// Async funtion ResponseAsisten
		var f = func() {
			go func() {
				sl.SystemGoPlay(ctx.system, parse)
				sl.RunCommand(parse)
			}()
		}
		// Return Response
		return parse(sl.ResponseNotParse), f
	}
	// Not Response Valide
	return "Lo siento, no puedo entenderte", func() {}
}

// WordsToSpeechBase64 convert words in speech base
func (ctx *AsistenCoreModel) WordsToSpeechBase64(words string) string {
	resp, err := ctx.clientts.ConvertToBase64(words)
	if err != nil {
		log.Println("ERROR=WordsToSpeechBase64(1) ", err)
	}
	return resp
}
