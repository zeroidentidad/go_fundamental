package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/olahol/melody"
)

//SetRealTimeRouter ruta para realtime de las acciones en la app
func SetRealTimeRouter(router *mux.Router) {
	melo := melody.New()
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		melo.HandleRequest(w, r)
	})

	melo.HandleMessage(func(s *melody.Session, msg []byte) {
		melo.Broadcast(msg)
	})
}
