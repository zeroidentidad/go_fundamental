package client

import (
	"fmt"
	"net/http"
)

// Controller incrusta un http.Client y lo usa internamente
type Controller struct {
	*http.Client
}

// DoOps (hacer operaciones) con un objeto controlador
func (c *Controller) DoOps() error {
	resp, err := c.Client.Get("http://www.google.com")
	if err != nil {
		return err
	}
	fmt.Println("results client.DoOps", resp.StatusCode)
	return nil
}
