package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// ServeHTTP ejemplo de handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Cualquier código que se desee en esta función")
}

func main() {

}
