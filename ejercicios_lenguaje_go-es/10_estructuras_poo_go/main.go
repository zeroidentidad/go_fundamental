package main

import (
	"fmt"
	"time"

	"./user"
)

/*type usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
	Categoria string
}*/

type jesus struct {
	user.Usuario
}

func main() {
	u := new(jesus)
	u.AltaUsuario(1, "Zero", time.Now(), true, "Admin")
	fmt.Println(u.Usuario)
}
