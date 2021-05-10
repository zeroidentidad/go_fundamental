package app

import (
	"os"
)

func Start() {
	configLoad()

	ADDR := os.Getenv("SERVER_ADDRESS")
	PORT := os.Getenv("SERVER_PORT")
	serve(ADDR, PORT, routes())
}
