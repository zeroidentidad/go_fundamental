package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zeroidentidad/hex_arq/internal/infrastructure/server"
)

func main() {
	engine := gin.Default()
	server.RegisterRouter("memory", engine)
	log.Fatal(engine.Run(":8080"))
}
