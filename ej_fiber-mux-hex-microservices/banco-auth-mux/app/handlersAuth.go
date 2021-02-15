package app

import (
	"encoding/json"
	"net/http"

	"github.com/zeroidentidad/fiber-hex-api-auth/dto"
	"github.com/zeroidentidad/fiber-hex-api-auth/service"
)

type HandlerAuth struct {
	service service.ServiceAuth
}

func (ha HandlerAuth) NotImplemented(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "Not implemented...")
}

func (ha HandlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	var requestLogin dto.RequestLogin
	if err := json.NewDecoder(r.Body).Decode(&requestLogin); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		token, err := ha.service.Login(requestLogin)
		if err != nil {
			writeResponse(w, http.StatusUnauthorized, err.Error())
		} else {
			writeResponse(w, http.StatusOK, *token)
		}
	}
}

/* .../auth/verify?token=validtokenstring&route=GetCliente&cliente_id=2000&cuenta_id=95470 */
func (ha HandlerAuth) Verify(w http.ResponseWriter, r *http.Request) {
	urlParams := make(map[string]string)

	// convertir de Query a tipo de mapa
	for k := range r.URL.Query() {
		urlParams[k] = r.URL.Query().Get(k)
	}

	if urlParams["token"] != "" {
		isAuthorized, appError := ha.service.Verify(urlParams)
		if appError != nil {
			writeResponse(w, http.StatusForbidden, notAuthorizedResponse())
		} else {
			if isAuthorized {
				writeResponse(w, http.StatusOK, authorizedResponse())
			} else {
				writeResponse(w, http.StatusForbidden, notAuthorizedResponse())
			}
		}
	} else {
		writeResponse(w, http.StatusForbidden, "missing token")
	}
}
