package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/zeroidentidad/fiber-hex-api-auth/domain"
	logs "github.com/zeroidentidad/fiber-hex-api-auth/logger"
	"github.com/zeroidentidad/fiber-hex-api-auth/service"
)

func Start() {
	sanityCheck()
	router := mux.NewRouter()

	storageDbAuth := domain.NewStorageDbAuth(getDbClient())
	hauth := HandlerAuth{service.NewServiceAuth(storageDbAuth, domain.GetRolePermissions())}

	router.HandleFunc("/auth/login", hauth.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", hauth.NotImplemented).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", hauth.Verify).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logs.Info(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
