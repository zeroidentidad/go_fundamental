package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zeroidentidad/servfiles/config"
)

var confObj, portString = *config.GetConfiguration(), fmt.Sprintf("%s", confObj.HttpPort)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	config.GetIPs()
	fmt.Println("---------------------\nPuerto config: " + portString + "\n---------------------")
	log.Fatal(http.ListenAndServe(":"+portString, nil))
}
