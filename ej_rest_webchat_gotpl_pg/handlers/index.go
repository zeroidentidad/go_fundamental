package handlers

import (
	"net/http"

	"webchat/db"
	"webchat/helper"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	threads, err := db.Threads()
	if err != nil {
		helper.ErrorMessage(writer, request, "Cannot get threads")
	} else {
		_, err := helper.Session(writer, request)
		if err != nil {
			helper.GenerateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			helper.GenerateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
