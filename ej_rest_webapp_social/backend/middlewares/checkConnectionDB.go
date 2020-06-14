package middlewares

import (
	"net/http"

	"github.com/zeroidentidad/twittor/db"
)

/*CheckConnectionDB permite conocer el estado de DB */
func CheckConnectionDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Conexión perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
