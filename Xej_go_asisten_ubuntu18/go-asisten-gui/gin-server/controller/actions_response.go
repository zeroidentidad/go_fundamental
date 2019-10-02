package controller

import (
	"fmt"
	"net/http"

	"../../mod-asisten-core/models"
	"github.com/gin-gonic/gin"
)

// ActionsResponseByActionAllHTML func
func ActionsResponseByActionAllHTML(ctx *gin.Context) {
	var actionsr = new([]models.ActionResponse)
	var action = new(models.Action)
	actionID := ctx.Query("action_id")
	Db.Find(actionsr, "action_id = ?", actionID)
	Db.Find(action, actionID)
	ctx.HTML(200, "actions_response_all.html", gin.H{
		"Elements": actionsr,
		"Action":   action,
	})
}

// ActionsResponseAddAPI func
func ActionsResponseAddAPI(ctx *gin.Context) {
	var actionr = new(models.ActionResponse)
	ctx.BindQuery(actionr)
	if actionr.Response == "" || actionr.ActionID == 0 {
		ctx.Redirect(http.StatusTemporaryRedirect, "/config/actions")
		return
	}
	Db.Create(actionr)
	ctx.Redirect(http.StatusTemporaryRedirect, "/config/actions_response?action_id="+fmt.Sprint(actionr.ActionID))
}

// ActionsResponseDeleteAPI func
func ActionsResponseDeleteAPI(ctx *gin.Context) {
	var actionr = new(models.ActionResponse)
	Db.Delete(actionr, ctx.Param("id"))
	ctx.Redirect(http.StatusTemporaryRedirect, "/config/actions_response?action_id="+ctx.Query("action_id"))
}
