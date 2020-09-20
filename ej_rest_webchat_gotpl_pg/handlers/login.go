package handlers

import (
	"net/http"
	"webchat/helper"
)

// GET /login
func Login(writer http.ResponseWriter, request *http.Request) {
	t := helper.ParseTemplateFiles("login.layout", "public.navbar", "login")
	_ = t.Execute(writer, nil)
}
