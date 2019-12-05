package handlers

import (
	"net/http"

	"../utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "app/index", nil)
}
