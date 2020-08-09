package errors

import "github.com/gin-gonic/gin"

// ValidationErrors valida errores entradas de usuario
var ValidationErrors = []string{}

//HandleErr manejador de errores genérico, logs de error y Os.Exit(1)
func HandleErr(c *gin.Context, err error) error {
	if err != nil {
		c.Error(err)
	}
	return err
}
