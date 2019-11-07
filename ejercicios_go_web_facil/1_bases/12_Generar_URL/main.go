package main

import (
	"fmt"
	"net/url"
)

func createURL() string {
	u, err := url.Parse("/params")

	if err != nil {
		panic(err)
	}

	u.Host = "localhost:8080"
	u.Scheme = "http"

	query := u.Query()
	query.Add("nombre", "valor")

	u.RawQuery = query.Encode()

	return u.String()
}

func main() {
	url := createURL()

	fmt.Println("URL formada: ", url)
}
