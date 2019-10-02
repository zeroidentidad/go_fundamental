package readconfig

import (
	"log"
	"math/rand"
	"strings"

	"github.com/jinzhu/gorm"
	//"../mod-asisten-core/models/methods"
	"../mod-asisten-core/models"
	"../osfunc"
	"../systemgo"
)

// ModelSpeechGroup Db
type ModelSpeechGroup struct {
	Db *gorm.DB
}

// ModelGroup models
type ModelGroup struct {
	valor            models.GroupKey
	SystemGo         string
	ResponseNotParse string
	Command          string
}

// NewSpeech methods
func NewSpeech(db *gorm.DB) *ModelSpeechGroup {
	var vl = new(ModelSpeechGroup)
	vl.Db = db
	return vl
}

// Find words in ModelGroup
func (ctx *ModelSpeechGroup) Find(words string) *ModelGroup {
	// Create Model Concidencias
	type ModelConcidencias struct {
		GroupKeyPosition int
		NumberConc       int
	}
	// Create Var Save Concidencias
	var arrConcidencias []*ModelConcidencias

	// Find GroupKeys
	var grpKey = models.GetAllGrupKeys(ctx.Db)
	for id, val := range grpKey {
		valid, conc := val.GroupKeyValidWords(words)
		if valid == true {
			arrConcidencias = append(arrConcidencias, &ModelConcidencias{
				GroupKeyPosition: id,
				NumberConc:       conc,
			})
		}
	}

	// Verificate Results
	if len(arrConcidencias) == 0 {
		return nil
	}

	// Busca el elemento con mayor concidencias
	var idMayor int
	var concMayor int
	for id, val := range arrConcidencias {
		if val.NumberConc > concMayor {
			concMayor = val.NumberConc
			idMayor = id
		}
	}

	// Prepare Values to return
	var vl = new(ModelGroup)
	vl.valor = grpKey[arrConcidencias[idMayor].GroupKeyPosition]

	var action = new(models.Action)
	ctx.Db.Find(action, vl.valor.ActionID)

	var actionResponse []models.ActionResponse
	ctx.Db.Find(&actionResponse, "action_id = ?", action.ID)
	if len(actionResponse) == 0 {
		log.Println("ACTION: Response not Avalibles")
	} else {
		vl.ResponseNotParse = actionResponse[rand.Int63n(int64(len(actionResponse)))].Response
	}

	// action
	vl.SystemGo = action.SystemGo
	vl.Command = action.Command

	// Return Response
	return vl
}

// SystemGoPlay run SystemGo Comand
func (ctx *ModelGroup) SystemGoPlay(cl *systemgo.SystemGoModel, parser func(string) string) error {
	if ctx.SystemGo == "" {
		return nil
	}
	ctx.SystemGo = parser(ctx.SystemGo)
	nameAndParams := strings.Split(ctx.SystemGo, ":")
	params := strings.Split(nameAndParams[1], ",")
	return cl.FindRun(nameAndParams[0], params...)
}

// RunCommand run comand
func (ctx *ModelGroup) RunCommand(parser func(string) string) error {
	if ctx.Command == "" {
		return nil
	}
	return osfunc.CommandRun(parser(ctx.Command))
}
