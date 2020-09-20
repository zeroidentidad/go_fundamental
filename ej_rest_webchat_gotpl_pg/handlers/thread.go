package handlers

import (
	"fmt"
	"net/http"
	"webchat/db"
	"webchat/helper"
	"webchat/utils"
)

// GET /threads/new
// mostrar página de hilo nuevo
func NewThread(writer http.ResponseWriter, request *http.Request) {
	_, err := helper.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		helper.GenerateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /thread/create
// crear hilo de usuario
func CreateThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := helper.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			utils.Danger(err, "No se puede procesar formulario")
		}
		user, err := sess.User()
		if err != nil {
			utils.Danger(err, "No se puede obtener usuario de sesión")
		}
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			utils.Danger(err, "No se puede crear hilo")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /thread/read
// mostrar detalles del hilo, incluidas las publicaciones
// y el formulario para escribir una publicación
func ReadThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := db.ThreadByUUID(uuid)
	if err != nil {
		helper.ErrorMessage(writer, request, "No se puede leer el hilo")
	} else {
		_, err := helper.Session(writer, request)
		if err != nil {
			helper.GenerateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
		} else {
			helper.GenerateHTML(writer, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

// POST /thread/post
// crear publicación
func PostThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := helper.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			utils.Danger(err, "No se puede procesar formulario")
		}
		user, err := sess.User()
		if err != nil {
			utils.Danger(err, "No se puede obtener usuario de sesión")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := db.ThreadByUUID(uuid)
		if err != nil {
			helper.ErrorMessage(writer, request, "No se puede leer el hilo")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			utils.Danger(err, "No se puede crear publicación")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
