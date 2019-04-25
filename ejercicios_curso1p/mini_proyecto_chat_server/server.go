
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "URL.Path = %q\n", r.URL.Path)
}