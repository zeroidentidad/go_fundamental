package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// ServeHTTP ejemplo de handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Cualquier código que se desee en esta función")
}

func main() {

	var d hotdog

	//Serve handler
	http.ListenAndServe(":8080", d)
}
