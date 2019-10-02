package controller

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"../../mod-asisten-core/models"
	"github.com/gin-gonic/gin"
)

// ActionsAddHTML func
func ActionsAddHTML(ctx *gin.Context) {
	ctx.HTML(200, "actions_add.html", gin.H{})
}

// ActionsAllHTML func
func ActionsAllHTML(ctx *gin.Context) {
	var actions = new([]models.Action)
	Db.Find(actions)
	ctx.HTML(200, "actions_all.html", gin.H{
		"Actions": actions,
	})
}

// ActionsAddBasicAPI methods
func ActionsAddBasicAPI(ctx *gin.Context) {
	var data = new(models.Action)
	ctx.BindQuery(data)
	responses := ctx.Query("Responses")
	responseArray := strings.Split(responses, ",")
	if data.Name == "" {
		log.Println("ERROR: Action name is null")
		ctx.Redirect(http.StatusTemporaryRedirect, "/config/actions/add")
		return
	}
	Db.Create(data)
	for _, val := range responseArray {
		if val != "" && data.ID != 0 {
			Db.Create(&models.ActionResponse{
				ActionID: data.ID,
				Response: val,
			})
		}
	}
	fmt.Printf("%v,%v\n", responses, responseArray)
	ctx.Redirect(http.StatusTemporaryRedirect, "/config/actions")
}
