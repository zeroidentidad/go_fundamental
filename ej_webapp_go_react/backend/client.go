package main

import (
	"log"

	r "github.com/dancannon/gorethink"
	"github.com/gorilla/websocket"
)

//FindHandler export in main
type FindHandler func(string) (Handler, bool)

//Client export in main
type Client struct {
	send          chan Mensaje
	socket        *websocket.Conn
	findHandler   FindHandler
	session       *r.Session
	stopCanales   map[int]chan bool
	id            string
	nombreUsuario string
}

//NewStopCanal export in main
func (c *Client) NewStopCanal(stopKey int) chan bool {
	c.StopForKey(stopKey)
	stop := make(chan bool)
	c.stopCanales[stopKey] = stop
	return stop
}

//StopForKey export in main
func (c *Client) StopForKey(key int) {
	if ch, found := c.stopCanales[key]; found {
		ch <- true
		delete(c.stopCanales, key)
	}
}

//Read export in main
func (c *Client) Read() {
	var mensaje Mensaje
	for {
		if err := c.socket.ReadJSON(&mensaje); err != nil {
			break
		}
		if handler, found := c.findHandler(mensaje.Nombre); found {
			handler(c, mensaje.Data)
		}
	}
	c.socket.Close()
}

//Write export in main
func (c *Client) Write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

//Close export in main
func (c *Client) Close() {
	for _, ch := range c.stopCanales {
		ch <- true
	}
	close(c.send)
	// delete usuario
	r.Table("usuario").Get(c.id).Delete().Exec(c.session)
}

//NewClient export in main
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
