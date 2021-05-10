package app

import "log"

func Start() {
	configLoad()

	db := dbClient()
	log.Println(db)

	serve(routes())
}
