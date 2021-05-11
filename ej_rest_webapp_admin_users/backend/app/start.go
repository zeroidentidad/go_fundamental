package app

import "log"

func Start() {
	config()

	db := dbclient()
	log.Println(db)

	serve(routes())
}
