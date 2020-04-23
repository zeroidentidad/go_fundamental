package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zeroidentidad/backend/database"
)

func main() {

	dbConn := database.InitDB()
	defer dbConn.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
		fmt.Println(dbConn)
	})
	http.ListenAndServe(":3000", r)
}
