package main

import (
	"time"

	r "github.com/dancannon/gorethink"
	"github.com/mitchellh/mapstructure"
)

const (
	CanalStop = iota
	UsuarioStop
	MensajeStop
)

//Mensaje export in main
type Mensaje struct {
	Nombre string      `json:"nombre"`
	Data   interface{} `json:"data"`
}

//Canal export in main
type Canal struct {
	Id     string `json:"id" gorethink:"id,omitempty"`
	Nombre string `json:"nombre" gorethink:"nombre"`
}

//Usuario export in main
type Usuario struct {
	Id     string `gorethink:"id,omitempty"`
	Nombre string `gorethink:"nombre"`
}

//CanalMensaje export in main
type CanalMensaje struct {
	Id        string    `gorethink:"id,omitempty"`
	CanalId   string    `gorethink:"canalId"`
	Body      string    `gorethink:"body"`
	Autor     string    `gorethink:"autor"`
	CreatedAt time.Time `gorethink:"createdAt"`
}

func editarUsuario(client *Client, data interface{}) {
	var usuario Usuario
	err := mapstructure.Decode(data, &usuario)
	if err != nil {
		client.send <- Mensaje{"error", err.Error()}
		return
	}
	client.nombreUsuario = usuario.Nombre
	go func() {
		_, err := r.Table("usuario").
			Get(client.id).
			Update(usuario).
			RunWrite(client.session)
		if err != nil {
			client.send <- Mensaje{"error", err.Error()}
		}
	}()
}

func subscribirUsuario(client *Client, data interface{}) {
	go func() {
		stop := client.NewStopCanal(UsuarioStop)
		cursor, err := r.Table("usuario").
			Changes(r.ChangesOpts{IncludeInitial: true}).
			Run(client.session)

		if err != nil {
			client.send <- Mensaje{"error", err.Error()}
			return
		}
		changeFeedHelper(cursor, "usuario", client.send, stop)
	}()
}

func desubscribirUsuario(client *Client, data interface{}) {
	client.StopForKey(UsuarioStop)
}

func agregarMensaje(client *Client, data interface{}) {
	var canalMensaje CanalMensaje
	err := mapstructure.Decode(data, &canalMensaje)
	if err != nil {
		client.send <- Mensaje{"error", err.Error()}
	}
	go func() {
		canalMensaje.CreatedAt = time.Now()
		canalMensaje.Autor = client.nombreUsuario
		err := r.Table("mensaje").
			Insert(canalMensaje).
			Exec(client.session)
		if err != nil {
			client.send <- Mensaje{"error", err.Error()}
		}
	}()
}

func subscribirMensaje(client *Client, data interface{}) {
	go func() {
		eventData := data.(map[string]interface{})
		val, ok := eventData["canalId"]
		if !ok {
			return
		}
		canalId, ok := val.(string)
		if !ok {
			return
		}
		stop := client.NewStopCanal(MensajeStop)
		cursor, err := r.Table("mensaje").
			OrderBy(r.OrderByOpts{Index: r.Desc("createdAt")}).
			Filter(r.Row.Field("canalId").Eq(canalId)).
			Changes(r.ChangesOpts{IncludeInitial: true}).
			Run(client.session)

		if err != nil {
			client.send <- Mensaje{"error", err.Error()}
			return
		}
		changeFeedHelper(cursor, "mensaje", client.send, stop)
	}()
}

func desubscribirMensaje(client *Client, data interface{}) {
	client.StopForKey(MensajeStop)
}

func agregarCanal(client *Client, data interface{}) {
	var canal Canal
	err := mapstructure.Decode(data, &canal)
	if err != nil {
		client.send <- Mensaje{"error", err.Error()}
		return
	}
	go func() {
		err = r.Table("canal").
			Insert(canal).
			Exec(client.session)
		if err != nil {
			client.send <- Mensaje{"error", err.Error()}
		}
	}()
}

func subscribirCanal(client *Client, data interface{}) {
	go func() {
		stop := client.NewStopCanal(CanalStop)
		cursor, err := r.Table("canal").
			Changes(r.ChangesOpts{IncludeInitial: true}).
			Run(client.session)
		if err != nil {
			client.send <- Mensaje{"error", err.Error()}
			return
		}
		changeFeedHelper(cursor, "canal", client.send, stop)
	}()
}

func desubscribirCanal(client *Client, data interface{}) {
	client.StopForKey(CanalStop)
}

func changeFeedHelper(cursor *r.Cursor, changeEventNombre string,
	send chan<- Mensaje, stop <-chan bool) {
	change := make(chan r.ChangeResponse)
	cursor.Listen(change)
	for {
		eventNombre := ""
		var data interface{}
		select {
		case <-stop:
			cursor.Close()
			return
		case val := <-change:
			if val.NewValue != nil && val.OldValue == nil {
				eventNombre = "agregar " + changeEventNombre
				data = val.NewValue
			} else if val.NewValue == nil && val.OldValue != nil {
				eventNombre = "remover " + changeEventNombre
				data = val.OldValue
			} else if val.NewValue != nil && val.OldValue != nil {
				eventNombre = "editar " + changeEventNombre
				data = val.NewValue
			}
			send <- Mensaje{eventNombre, data}
		}
	}
}
