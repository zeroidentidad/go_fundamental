package app

import (
	"backend/logs"
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func serve(router *fiber.App) {
	addr := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	go func() {
		logs.Info(fmt.Sprintf("Starting server on %s:%s ...", addr, port))
		err := router.Listen(fmt.Sprintf("%s:%s", addr, port))
		logs.Fatal(err.Error())
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
