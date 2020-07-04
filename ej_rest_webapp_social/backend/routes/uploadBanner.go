package routes

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/zeroidentidad/twittor/db"
	"github.com/zeroidentidad/twittor/models"
)

/*UploadBanner envia subida de Banner al servidor*/
func UploadBanner(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banners/" + IDUser + "." + extension

	of, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(of, file)
	if err != nil {
		http.Error(w, "Error al copiar imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Banner = IDUser + "." + extension
	status, err = db.ModifyRegistry(user, IDUser)
	if err != nil || status == false {
		http.Error(w, "Error al guardar banner en DB! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
