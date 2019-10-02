package controller

import (
	"log"
	"net/http"
	"strings"

	"../../mod-asisten-core/models"
	"github.com/gin-gonic/gin"
)

// ProgramsAllHTML func
func ProgramsAllHTML(ctx *gin.Context) {
	var programs []models.ProgramsAlternative
	Db.Find(&programs)
	for i, j := 0, len(programs)-1; i < j; i, j = i+1, j-1 {
		programs[i], programs[j] = programs[j], programs[i]
	}
	ctx.HTML(200, "programs_all.html", gin.H{
		"Elements": programs,
	})
}

// ProgramsDeleteAPI func
func ProgramsDeleteAPI(ctx *gin.Context) {
	id := ctx.Param("id")
	Db.Delete(&models.ProgramsAlternative{}, id)
	ctx.Redirect(http.StatusTemporaryRedirect, "/config/programs")
}

// ProgramsAddAPI func
func ProgramsAddAPI(ctx *gin.Context) {
	programsString, _ := ctx.GetQuery("Programas")
	programs := strings.Split(programsString, ",")
	generalName, _ := ctx.GetQuery("GeneralName")
	if len(programs) == 0 {
		log.Println("ERROR: Not Programs")
		ctx.Redirect(http.StatusTemporaryRedirect, "/config/programs")
		return
	}
	if generalName == "" {
		log.Println("ERROR: Not GeneralName")
		ctx.Redirect(http.StatusTemporaryRedirect, "/config/programs")
		return
	}
	var pro = new(models.ProgramsAlternative)
	pro.Programs = strings.Join(programs, ":")
	pro.GeneralName = strings.ToUpper(generalName)
	Db.Create(pro)
	ctx.Redirect(http.StatusTemporaryRedirect, "/config/programs")
}
