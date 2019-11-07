package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", watahea)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func watahea(w http.ResponseWriter, req *http.Request) {

	var s, mimetype string
	if req.Method == http.MethodPost {

		// abrir
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// sobre la info en terminal
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		// leer buffer string
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)

		// almacenar en el servidor
		dst, err := os.Create(filepath.Join("./userfiles/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Printf("Extension archivo: %q\n", filepath.Ext(h.Filename))

		if filepath.Ext(h.Filename) == ".txt" {
			mimetype = "text/html; charset=utf-8"
		} else if filepath.Ext(h.Filename) == ".png" {
			mimetype = "image/png;"
		}
	}

	w.Header().Set("Content-Type", mimetype)
	tpl.ExecuteTemplate(w, "index.gohtml", s)
}
