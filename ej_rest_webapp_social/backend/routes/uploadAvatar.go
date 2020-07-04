package routes

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/zeroidentidad/twittor/db"
	"github.com/zeroidentidad/twittor/models"
)

/*UploadAvatar envia subida de Avatar al servidor*/
func UploadAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IDUser + "." + extension

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

	user.Avatar = IDUser + "." + extension
	status, err = db.ModifyRegistry(user, IDUser)
	if err != nil || status == false {
		http.Error(w, "Error al guardar avatar en DB! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
