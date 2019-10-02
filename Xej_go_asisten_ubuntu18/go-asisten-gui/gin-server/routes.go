package server

import (
	"net/http"

	"./controller"
	"github.com/gin-gonic/gin"
)

// Routes app
func Routes(route *gin.Engine) {
	// Especial Routes
	route.GET("/asisten", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "asisten.html", nil)
	})
	route.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/config")
	})
	// Config Routes
	var config = route.Group("/config")
	{
		config.GET("/", func(ctx *gin.Context) {
			ctx.HTML(200, "home.html", nil)
		})
		config.GET("/actions", controller.ActionsAllHTML)
		config.GET("/actions/add", controller.ActionsAddHTML)
		config.GET("/actions/add-api", controller.ActionsAddBasicAPI)
		config.GET("/actions_response", controller.ActionsResponseByActionAllHTML)
		config.GET("/actions_response/add", controller.ActionsResponseAddAPI)
		config.GET("/actions_response/delete/:id", controller.ActionsResponseDeleteAPI)
		config.GET("/command_keys", controller.CommandKeysAllHTML)
		config.GET("/command_keys/add", controller.CommandKeysAddAPI)
		config.GET("/command_keys/delete/:id", controller.CommandKeysDeleteAPI)
		config.GET("/keywords", controller.KeyWordsAllHTML)
		config.GET("/keywords/delete/:id", controller.KeyWordsDeleteAPI)
		config.GET("/keywords/add", controller.KeyWordsAddAPI)
		config.GET("/programs", controller.ProgramsAllHTML)
		config.GET("/programs/add", controller.ProgramsAddAPI)
		config.GET("/programs/delete/:id", controller.ProgramsDeleteAPI)
		config.GET("/group_keys", controller.GroupKeysAllHTML)
		config.GET("/group_keys/add-api", controller.GroupKeysAddAPI)
		config.GET("/group_keys/delete/:id", controller.GroupKeysDeleteAPI)
	}
}
