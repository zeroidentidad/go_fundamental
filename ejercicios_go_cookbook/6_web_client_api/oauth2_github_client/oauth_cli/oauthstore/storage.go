package oauthstore

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

// Storage es la interfaz de almacenamiento genérica
type Storage interface {
	GetToken() (*oauth2.Token, error)
	SetToken(*oauth2.Token) error
}

// GetToken recupera token de github oauth2
func GetToken(ctx context.Context, conf Config) (*oauth2.Token, error) {
	token, err := conf.Storage.GetToken()
	if err == nil && token.Valid() {
		return token, err
	}
	url := conf.AuthCodeURL("state")
	fmt.Printf("Ingresa la URL en navegador web y siga instrucciones en pantalla: %v\n", url)
	fmt.Println("Pegue código devuelto en URL de redireccionamiento y presione Enter:")

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, err
	}
	return conf.Exchange(ctx, code)
}
