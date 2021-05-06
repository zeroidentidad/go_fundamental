package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	CONN_PORT = "8080"
)

var newCache *cache.Cache

func init() {
	newCache = cache.New(5*time.Minute, 10*time.Minute)
	newCache.Set("gophers", "gophers latam", cache.DefaultExpiration)
}

func getFromCache(w http.ResponseWriter, r *http.Request) {
	gophers, found := newCache.Get("gophers")

	if found {
		log.Print("Key encontrada en caché con valor, como : ", gophers.(string))
		fmt.Fprintf(w, "Hola "+gophers.(string))
	} else {
		log.Print("Key no encontrada en caché : ", "gophers")
		fmt.Fprintf(w, "Key no encontrada en caché")
	}
}

func main() {
	http.HandleFunc("/", getFromCache)

	err := http.ListenAndServe(":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error al iniciar el servidor http : ", err)
		return
	}
}
