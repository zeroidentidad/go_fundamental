package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/zeroidentidad/servfiles/config"
)

func main() {
	config.GetIPs()
	config.PortHttp(true)
	http.Handle("/", config.NoCache(http.FileServer(http.Dir(filepath.Dir(config.GetDir())))))
	config.UrlExample()
	fmt.Println("\nCtrl+C para salir o cerrar ventana...")
	log.Fatal(http.ListenAndServe(":"+config.PortHttp(false), nil))
}
