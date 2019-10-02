package server

import (
	"./controller"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// InitServer cargar server
func InitServer(port string, db *gorm.DB, dirAssets string) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	controller.Db = db
	app.LoadHTMLGlob(dirAssets + "templates/*")
	app.Static("/assets", dirAssets+"assets")
	Routes(app)
	go func() {
		app.Run(":" + port)
	}()
}
