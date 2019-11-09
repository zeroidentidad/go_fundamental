package main

import (
	"fmt"
	"log"
	"net/http"
)

type xHandler func(http.ResponseWriter, *http.Request)

type muxPersonal struct {
	rutasPersonales map[string]xHandler // asoc handlers
}

func (este *muxPersonal) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn := este.rutasPersonales[r.URL.Path]
	fn(w, r)
}

func (este *muxPersonal) AddMux(ruta string, fn xHandler) {
	este.rutasPersonales[ruta] = fn
}

func main() {

	newMap := make(map[string]xHandler)
	mux := &muxPersonal{rutasPersonales: newMap}
	mux.AddMux("/holas", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "watahea from func anonima")
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
