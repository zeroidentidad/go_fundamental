package main

import (
	"context"

	"oauth-cli/oauthconf"
)

func main() {
	ctx := context.Background()
	conf := oauthconf.Setup()

	tok, err := oauthconf.GetToken(ctx, conf)
	if err != nil {
		panic(err)
	}
	client := conf.Client(ctx, tok)

	if err := oauthconf.GetUser(client); err != nil {
		panic(err)
	}

}
