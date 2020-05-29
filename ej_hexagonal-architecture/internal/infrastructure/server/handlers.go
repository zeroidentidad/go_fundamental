package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zeroidentidad/hex_arq/internal/core/domain"
	"github.com/zeroidentidad/hex_arq/internal/core/ports"
)

func GetPersonEndpoint(service ports.PersonsService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var person domain.Person
		id, isExist := c.Params.Get("id")
		if isExist {
			idNumb, err := strconv.Atoi(id)
			if err != nil {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			person, err = service.Get(idNumb)
			if err != nil {
				c.AbortWithStatus(http.StatusNoContent)
				return
			}
		}
		c.JSON(http.StatusOK, BuildResponsePersonGet(person))
	}
}

func GetUserEndpoint(service ports.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user domain.User
		id, isExist := c.Params.Get("id")
		if isExist {
			_, err := strconv.Atoi(id)
			if err != nil {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			user = service.Get()
			if err != nil {
				c.AbortWithStatus(http.StatusNoContent)
				return
			}
		}
		c.JSON(http.StatusOK, BuildResponseUserGet(user))
	}
}
