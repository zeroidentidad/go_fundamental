package client

import (
	"fmt"
	"net/http"
)

// DoOps toma un cliente, luego busca facebook.com
func DoOps(c *http.Client) error {
	resp, err := c.Get("http://www.facebook.com")
	if err != nil {
		return err
	}
	fmt.Println("results DoOps:", resp.StatusCode)

	return nil
}

// DefaultGetGolang usa el cliente predeterminado
// para obtener golang.org
func DefaultGetGolang() error {
	resp, err := http.Get("https://www.golang.org")
	if err != nil {
		return err
	}
	fmt.Println("results DefaultGetGolang:", resp.StatusCode)
	return nil
}
