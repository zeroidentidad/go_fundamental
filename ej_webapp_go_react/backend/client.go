package main

import (
	"log"

	r "github.com/dancannon/gorethink"
	"github.com/gorilla/websocket"
)

type FindHandler func(string) (Handler, bool)

type Client struct {
	send          chan Mensaje
	socket        *websocket.Conn
	findHandler   FindHandler
	session       *r.Session
	stopCanales   map[int]chan bool
	id            string
	nombreUsuario string
}

func (c *Client) NewStopChannel(stopKey int) chan bool {
	c.StopForKey(stopKey)
	stop := make(chan bool)
	c.stopCanales[stopKey] = stop
	return stop
}

func (c *Client) StopForKey(key int) {
	if ch, found := c.stopCanales[key]; found {
		ch <- true
		delete(c.stopCanales, key)
	}
}

func (client *Client) Read() {
	var mensaje Mensaje
	for {
		if err := client.socket.ReadJSON(&mensaje); err != nil {
			break
		}
		if handler, found := client.findHandler(mensaje.Nombre); found {
			handler(client, mensaje.Data)
		}
	}
	client.socket.Close()
}

func (client *Client) Write() {
	for msg := range client.send {
		if err := client.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	client.socket.Close()
}

func (c *Client) Close() {
	for _, ch := range c.stopCanales {
		ch <- true
	}
	close(c.send)
	// delete usuario
	r.Table("usuario").Get(c.id).Delete().Exec(c.session)
}

func NewClient(socket *websocket.Conn, findHandler FindHandler,
	session *r.Session) *Client {
	var usuario Usuario
	usuario.Nombre = "anonimo"
	res, err := r.Table("usuario").Insert(usuario).RunWrite(session)
	if err != nil {
		log.Println(err.Error())
	}
	var id string
	if len(res.GeneratedKeys) > 0 {
		id = res.GeneratedKeys[0]
	}
	return &Client{
		send:          make(chan Mensaje),
		socket:        socket,
		findHandler:   findHandler,
		session:       session,
		stopCanales:   make(map[int]chan bool),
		id:            id,
		nombreUsuario: usuario.Nombre,
	}
}
