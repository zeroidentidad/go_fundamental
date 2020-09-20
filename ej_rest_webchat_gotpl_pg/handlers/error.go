package handlers

import (
	"net/http"
	"webchat/helper"
)

// GET /err?msg=
// mostrar página del mensaje de error
func Err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := helper.Session(writer, request)
	if err != nil {
		helper.GenerateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		helper.GenerateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}
