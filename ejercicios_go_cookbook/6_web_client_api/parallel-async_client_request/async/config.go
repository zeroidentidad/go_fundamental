package async

import "net/http"

// NewClient crea un nuevo cliente y establece sus canales apropiados
func NewClient(client *http.Client, bufferSize int) *Client {
	respch := make(chan *http.Response, bufferSize)
	errch := make(chan error, bufferSize)
	return &Client{
		Client: client,
		Resp:   respch,
		Err:    errch,
	}
}

// Client contiene un cliente y tiene dos canales para agregar
// respuestas y errores
type Client struct {
	*http.Client
	Resp chan *http.Response
	Err  chan error
}

// AsyncGet realiza una peticion GET y luego devuelve
// la respuesta/error al canal apropiado
func (c *Client) AsyncGet(url string) {
	resp, err := c.Get(url)
	if err != nil {
		c.Err <- err
		return
	}
	c.Resp <- resp
}
