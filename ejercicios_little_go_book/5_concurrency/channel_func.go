package main

import (
	"fmt"
	"time"
)

type Usuario struct {
	Name string
	Edad int
}

var usuario Usuario

func main() {
	usuario = Usuario{Name: "Guillermo romo", Edad: 28}
	usuarioChanel := make(chan Usuario)

	go usuario.enviarUsuario(usuarioChanel)
	time.Sleep(time.Second * 5)
	go usuario.RecibirUsuario(usuarioChanel)
	time.Sleep(time.Second * 5)
}

func (usuario Usuario) enviarUsuario(usuarioChanel chan Usuario) {
	for {
		usuarioChanel <- usuario
	}
}

func (usuario Usuario) RecibirUsuario(usuarioChanel chan Usuario) {
	for {
		fmt.Printf("T", usuarioChanel)
		usuario := <-usuarioChanel
		fmt.Println(usuario)
		time.Sleep(time.Second * 1)
	}
}
