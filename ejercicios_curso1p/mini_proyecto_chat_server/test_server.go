
package main

import (
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	http.Redirect(w, r, "http://www.google.com", 301)
	// http.Redirect(w, r, "http://wwww.google.com", http.StatusMovedPermanently)
}