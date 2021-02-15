package app

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/zeroidentidad/fiber-hex-api/domain"
	logs "github.com/zeroidentidad/fiber-hex-api/logger"
	"github.com/zeroidentidad/fiber-hex-api/service"
)

func Start() {
	sanityCheck()
	router := fiber.New()
	router.Use(logger.New())

	dbClient := getDbClient()
	storageDbCliente := domain.NewStorageDbCliente(dbClient)
	storageDbCuenta := domain.NewStorageDbCuenta(dbClient)

	hclientes := HandlersCliente{service.NewServiceCliente(storageDbCliente)}
	hcuentas := HandlersCuenta{service.NewServiceCuenta(storageDbCuenta)}

	router.Get("/clientes", hclientes.getAllClientes)
	router.Get("/clientes/:id", hclientes.getCliente)
	router.Post("/clientes/:id/cuenta", hcuentas.PostNewCuenta)
	router.Post("/clientes/:id/cuenta/:id_cuenta", hcuentas.PostNewTransaccion)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logs.Info(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	err := router.Listen(fmt.Sprintf("%s:%s", address, port))
	logs.Fatal(err.Error())
}

func sanityCheck() {
	err := godotenv.Load()
	if err != nil {
		logs.Fatal("Error loading .env file")
	}

	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logs.Fatal(fmt.Sprintf("- %s not defined. Terminating app...", k))
		}
	}
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	//pool settings:
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
