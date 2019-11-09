package mux

import (
	"net/http"
)

type xHandler func(http.ResponseWriter, *http.Request)

type muxPersonal struct {
	rutasPersonales map[string]xHandler // asoc handlers
}

func (este *muxPersonal) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if fn, ok := este.rutasPersonales[r.URL.Path]; ok {
		fn(w, r)
	} else {
		http.NotFound(w, r)
	}
	//fn := este.rutasPersonales[r.URL.Path]
}

func (este *muxPersonal) AddFunc(ruta string, fn xHandler) {
	este.rutasPersonales[ruta] = fn
}

func (este *muxPersonal) AddHandle(ruta string, hd http.Handler) {
	este.rutasPersonales[ruta] = hd.ServeHTTP
}

func CreateMux() *muxPersonal {
	newMap := make(map[string]xHandler)
	return &muxPersonal{newMap}
}
