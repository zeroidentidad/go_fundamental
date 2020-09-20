package main

import (
	"net/http"
	"time"

	"webchat/handlers"
	"webchat/utils"
)

func main() {

	utils.P("WebChat", utils.Version(), "started at", utils.Config.Address)

	mux := http.NewServeMux()

	files := http.FileServer(http.Dir(utils.Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// handlers funcs
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/err", handlers.Err)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/logout", handlers.Logout)
	mux.HandleFunc("/signup", handlers.Signup)
	mux.HandleFunc("/signup_account", handlers.SignupAccount)
	mux.HandleFunc("/authenticate", handlers.Authenticate)
	mux.HandleFunc("/thread/new", handlers.NewThread)
	mux.HandleFunc("/thread/create", handlers.CreateThread)
	mux.HandleFunc("/thread/post", handlers.PostThread)
	mux.HandleFunc("/thread/read", handlers.ReadThread)

	server := &http.Server{
		Addr:           utils.Config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(utils.Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(utils.Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	_ = server.ListenAndServe()
}
