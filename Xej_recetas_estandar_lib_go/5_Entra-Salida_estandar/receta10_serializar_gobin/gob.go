package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type User struct {
	Nombre   string
	Apellido string
	Edad     int
	Activo   bool
}

func (u User) String() string {
	return fmt.Sprintf(`{"Nombre":%s,"Apellido":%s,"Edad":%d,"Activo":%v }`,
		u.Nombre, u.Apellido, u.Edad, u.Activo)
}

type SimpleUser struct {
	Nombre   string
	Apellido string
}

func (u SimpleUser) String() string {
	return fmt.Sprintf(`{"Nombre":%s,"Apellido":%s}`,
		u.Nombre, u.Apellido)
}

func main() {

	var buff bytes.Buffer

	// Encode value
	enc := gob.NewEncoder(&buff)
	user := User{
		"Fulano",
		"De Tal",
		30,
		true,
	}
	enc.Encode(user)
	fmt.Printf("%X\n", buff.Bytes())

	// Decode value
	out := User{}
	dec := gob.NewDecoder(&buff)
	dec.Decode(&out)
	fmt.Println(out.String())

	enc.Encode(user) // x2
	out2 := SimpleUser{}
	dec.Decode(&out2)
	fmt.Println(out2.String())

}
