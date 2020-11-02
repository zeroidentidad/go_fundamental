package oauthstore

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

// Config envuelve el oauth2.Config predeterminado y
// agrega el almacenamiento personalizado
type Config struct {
	*oauth2.Config
	Storage
}

// Exchange almacena token después de la recuperación
func (c *Config) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := c.Config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	if err := c.Storage.SetToken(token); err != nil {
		return nil, err
	}
	return token, nil
}

// A TokenSource se le puede pasar un token para almacenar,
// o cuando se recupera uno nuevo, se almacena
func (c *Config) TokenSource(ctx context.Context, t *oauth2.Token) oauth2.TokenSource {
	return StorageTokenSource(ctx, c, t)
}

// El cliente se adjunta al TokenSource
func (c *Config) Client(ctx context.Context, t *oauth2.Token) *http.Client {
	return oauth2.NewClient(ctx, c.TokenSource(ctx, t))
}
