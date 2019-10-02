package readconfig

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

// ModelAsisten export
type ModelAsisten struct {
	Lang      string
	Name      string
	SpeechKey string
}

// ReadAsisten func
func ReadAsisten(DataDir string) (*ModelAsisten, error) {
	var model = new(ModelAsisten)
	file, err := ioutil.ReadFile(DataDir + "config/asisten.json")
	if err != nil {
		return nil, err
	}
	return model, json.Unmarshal(file, model)
}

// Parse func
func (ctx *ModelAsisten) Parse(words string, speech string) string {
	return strings.NewReplacer(
		"{{NAMEASISTEN}}", ctx.Name,
		"{{"+ctx.SpeechKey+"}}", speech,
	).Replace(words)
}
