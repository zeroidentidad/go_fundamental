package handlers

import (
	"fmt"
	"net/http"

	"../utils"
)

type customeHandler func(w http.ResponseWriter, r *http.Request)

func Authentication(function customeHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !utils.IsAuthenticated(r) {
			http.Redirect(w, r, "/users/login", http.StatusSeeOther)
			return
		}
		function(w, r)
	})
}

func MiddlewareTwo(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//logica
		fmt.Println("Wrap template Handler")
		handler.ServeHTTP(w, r)
	})
}
