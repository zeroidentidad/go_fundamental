package oauthconf

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// Setup devuelve un oauth2Config configurado para comunicar con
// github, se necesitan las variables de entorno establecidas para
// su identificación y secreto
func Setup() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT"),
		ClientSecret: os.Getenv("GITHUB_SECRET"),
		Scopes:       []string{"repo", "user"},
		Endpoint:     github.Endpoint,
	}
}

// GetToken recupera un token de github oauth2
func GetToken(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error) {
	url := conf.AuthCodeURL("state")
	fmt.Printf("Ingresa la URL en navegador web y siga instrucciones en pantalla: %v\n", url)
	fmt.Println("Pegue código devuelto en URL de redireccionamiento y presione Enter:")

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, err
	}
	return conf.Exchange(ctx, code)
}
