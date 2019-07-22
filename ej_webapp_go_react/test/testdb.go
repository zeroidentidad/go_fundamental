package main

import (
	"fmt"

	r "github.com/dancannon/gorethink"
)

type Usuario struct {
	Id     string `gorethink:"id,omitempty"`
	Nombre string `gorethink:"nombre"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "192.168.0.101:28015",
		Database: "rtgoreact",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	/*usuario := Usuario{
		Nombre: "Jesus Ferrer",
	}*/

	/*respuesta, err := r.Table("usuario").Insert(usuario).RunWrite(session)

	if err != nil {
		fmt.Println(err)
		return
	}*/

	/*respuesta, _ := r.Table("usuario").Get("817bd110-2aac-4dd3-87c5-5cb7f2d92cdb").Update(usuario).RunWrite(session)

	fmt.Printf("%#v\n", respuesta)*/

	cursor, _ := r.Table("usuario").Changes(r.ChangesOpts{IncludeInitial: true}).Run(session)
	var respuestaCambios r.ChangeResponse

	for cursor.Next(&respuestaCambios) {
		fmt.Printf("%#v\n", respuestaCambios)
	}

}
