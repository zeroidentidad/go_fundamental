package domain

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"github.com/zeroidentidad/fiber-hex-api/logger"
)

type StorageAuth interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}

type RemoteStorageAuth struct {
}

func (r RemoteStorageAuth) IsAuthorized(token string, routeName string, vars map[string]string) bool {
	u := buildVerifyURL(token, routeName, vars)

	if response, err := http.Get(u); err != nil {
		return false
	} else {
		m := map[string]bool{}
		if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
			logger.Error("Error while decoding response from auth server:" + err.Error())
			return false
		}
		return m["isAuthorized"]
	}
}

/*
  URL para la verificación del token en el siguiente formato:
  /auth/verify?token={token string}
  &routeName={current route name}
  &cliente_id={cliente_id current route}
  &cuenta_id={cuenta_id current route de estar disponible}

  Sample: /auth/verify?token=aa.bb.cc&routeName=postNewTransaccion&cliente_id=2000&cuenta_id=95470
*/
func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	u := url.URL{Host: os.Getenv("REMOTE_SERVER"), Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func NewStorageAuth() RemoteStorageAuth {
	return RemoteStorageAuth{}
}
