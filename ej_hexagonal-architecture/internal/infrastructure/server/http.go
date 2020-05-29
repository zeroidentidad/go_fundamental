package server

import (
	"github.com/gin-gonic/gin"
	"github.com/zeroidentidad/hex_arq/internal/core/ports"
	"github.com/zeroidentidad/hex_arq/internal/core/services/personssrv"
	"github.com/zeroidentidad/hex_arq/internal/core/services/usersserv"
	"github.com/zeroidentidad/hex_arq/internal/infrastructure/repositories/memory"
)

func FactoryRepository(engineRepo string) ports.PersonsRepository {
	if engineRepo == "memory" {
		return memory.NewPersonsRepository()
	} else if engineRepo == "sql" {
		return memory.NewPersonsRepository()
	}

	return nil
}

func RegisterRouter(engineRepo string, engine *gin.Engine) {

	//repoPerson := memory.NewPersonsRepository()
	repoPerson := FactoryRepository(engineRepo)
	servPerson := personssrv.NewService(repoPerson)
	getPersonEndpoint := GetPersonEndpoint(servPerson)

	repoUser := memory.NewUsersRepository()
	servUser := usersserv.NewService(repoUser)
	getUserEndpoint := GetUserEndpoint(servUser)

	engine.GET("person/:id", getPersonEndpoint)
	engine.GET("user/:id", getUserEndpoint)

}
