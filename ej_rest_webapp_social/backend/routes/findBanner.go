package routes

import (
	"io"
	"net/http"
	"os"

	"github.com/zeroidentidad/twittor/db"
)

/*FindBanner envia Banner del servidor al HTTP */
func FindBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/banners/" + perfil.Banner)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
	}
}
