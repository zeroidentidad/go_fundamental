package controller

import (
	"log"
	"net/http"
	"strings"

	"../../mod-asisten-core/models"
	"github.com/gin-gonic/gin"
)

// CommandKeysAllHTML func
func CommandKeysAllHTML(ctx *gin.Context) {
	var commandKeys []models.CommandKey
	Db.Find(&commandKeys)
	for i, j := 0, len(commandKeys)-1; i < j; i, j = i+1, j-1 {
		commandKeys[i], commandKeys[j] = commandKeys[j], commandKeys[i]
	}
	ctx.HTML(200, "command_keys_all.html", gin.H{
		"CommandKeys": commandKeys,
	})
}

// CommandKeysAddAPI func
func CommandKeysAddAPI(ctx *gin.Context) {
	var commandKey = new(models.CommandKey)
	ctx.BindQuery(commandKey)
	if commandKey.Key == "" || commandKey.Command == "" {
		log.Println("ERROR: Faltan parametros en commandKey")
		ctx.Redirect(http.StatusTemporaryRedirect, "/config/command_keys")
		return
	}
	commandKey.Key = strings.ToUpper(commandKey.Key)
	Db.Create(commandKey)
	ctx.Redirect(http.StatusTemporaryRedirect, "/config/command_keys")
}

// CommandKeysDeleteAPI func
func CommandKeysDeleteAPI(ctx *gin.Context) {
	id := ctx.Param("id")
	Db.Delete(&models.CommandKey{}, id)
	ctx.Redirect(http.StatusTemporaryRedirect, "/config/command_keys")
}
