package router

import (
	"postgres-jwt-react/controller"
	"postgres-jwt-react/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura enrutamiento
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Middlewares
	router.Use(middlewares.ErrorHandler)
	router.Use(middlewares.CORSMiddleware())

	// rutas
	router.GET("/ping", controller.Pong)
	router.POST("/register", controller.Create)
	router.POST("/login", controller.Login)
	router.GET("/session", controller.Session)
	router.POST("/createReset", controller.InitiatePasswordReset)
	router.POST("/resetPassword", controller.ResetPassword)
	return router
}
