package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	// sockets ver. ant.
	/*server.On("connection", func(so socketio.Socket) {
		log.Println("Usuario conectado")
		so.Join("chat_room")
		so.On("mensaje chat", func(msg string) {
			log.Println("emit: ", so.Emit("mensaje chat", msg))
			so.BroadcastTo("chat_room", "mensaje chat", msg)
		})
	})*/

	// sockets ver. act.
	server.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("mensaje chat")
		log.Println("Usuario conectado:", so.ID())

		server.OnEvent("/", "mensaje chat", func(so socketio.Conn, msg string) string {
			fmt.Println("emit:", msg)
			so.SetContext("mensaje chat")
			so.Emit("mensaje chat", msg)
			return msg
		})

		return nil
	})

	go server.Serve()
	defer server.Close()

	//http server
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
