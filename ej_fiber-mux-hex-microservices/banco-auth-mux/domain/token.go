package domain

import (
	"encoding/json"

	"github.com/dgrijalva/jwt-go"
)

const HMAC_SECRET = "hmac-sample-secret-xyz"

type Claims struct {
	ClienteID string   `json:"cliente_id"`
	Cuentas   []string `json:"cuentas"`
	Username  string   `json:"username"`
	Expiry    int64    `json:"exp"`
	Rol       string   `json:"rol"`
}

func (c Claims) IsUserRole() bool {
	return c.Rol == "user"
}

func BuildClaimsFromJwtMapClaims(mapClaims jwt.MapClaims) (*Claims, error) {
	bytes, err := json.Marshal(mapClaims)
	if err != nil {
		return nil, err
	}
	var c Claims
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (c Claims) IsValidClienteID(clienteID string) bool {
	return c.ClienteID == clienteID
}

func (c Claims) IsValidCuentaID(cuentaID string) bool {
	if cuentaID != "" {
		cuentaEncontrada := false
		for _, a := range c.Cuentas {
			if a == cuentaID {
				cuentaEncontrada = true
				break
			}
		}
		return cuentaEncontrada
	}
	return true
}

func (c Claims) IsRequestVerifiedWithTokenClaims(urlParams map[string]string) bool {
	if c.ClienteID != urlParams["cliente_id"] {
		return false
	}

	if !c.IsValidCuentaID(urlParams["cuenta_id"]) {
		return false
	}
	return true
}
