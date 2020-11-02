package oauthstore

import (
	"context"

	"golang.org/x/oauth2"
)

type storageTokenSource struct {
	*Config
	oauth2.TokenSource
}

// Token para satisfacer la interfaz TokenSource
func (s *storageTokenSource) Token() (*oauth2.Token, error) {
	if token, err := s.Config.Storage.GetToken(); err == nil && token.Valid() {
		return token, err
	}
	token, err := s.TokenSource.Token()
	if err != nil {
		return token, err
	}
	if err := s.Config.Storage.SetToken(token); err != nil {
		return nil, err
	}
	return token, nil
}

// StorageTokenSource será utilizado por el método config.TokenSource
func StorageTokenSource(ctx context.Context, c *Config, t *oauth2.Token) oauth2.TokenSource {
	if t == nil || !t.Valid() {
		if tok, err := c.Storage.GetToken(); err == nil {
			t = tok
		}
	}
	ts := c.Config.TokenSource(ctx, t)
	return &storageTokenSource{c, ts}
}
