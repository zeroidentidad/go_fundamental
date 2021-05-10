package app

import (
	"backend/logs"
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func serve(router *fiber.App) {
	ADDR := os.Getenv("SERVER_ADDRESS")
	PORT := os.Getenv("SERVER_PORT")
	go func() {
		logs.Info(fmt.Sprintf("Starting server on %s:%s ...", ADDR, PORT))
		err := router.Listen(fmt.Sprintf("%s:%s", ADDR, PORT))
		logs.Fatal(err.Error())
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
