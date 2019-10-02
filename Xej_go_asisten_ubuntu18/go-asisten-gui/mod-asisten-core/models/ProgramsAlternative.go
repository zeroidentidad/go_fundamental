package models

import (
	"errors"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
)

// ProgramsAlternative model find
type ProgramsAlternative struct {
	gorm.Model
	GeneralName string
	Programs    string
}

// Exits program exits
func (ctx *ProgramsAlternative) Exits() (string, error) {
	programs := strings.Split(ctx.Programs, ":")
	for _, name := range programs {
		_, err := os.Open("/bin/" + name)
		if err == nil {
			return name, nil
		}
	}

	return "", errors.New("Not Program exits")
}
