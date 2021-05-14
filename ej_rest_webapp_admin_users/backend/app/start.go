package app

func Start() {
	config()
	serve(routes("/api"))
}
