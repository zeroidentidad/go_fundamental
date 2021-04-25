package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Client struct {
	consultaIP string
	connString string
}

func (c *Client) String() string {
	return fmt.Sprintf("ConsultaIP: %s, String conn: %s",
		c.consultaIP, c.connString)
}

//valores por defecto sin el config.json
var defaultClient = Client{
	consultaIP: "localhost:9000",
	connString: "postgres://localhost:5432",
}

// ConfigFunc funciona como un type para ser utilizado
// en opciones funcionales
type ConfigFunc func(opt *Client)

// FromFile func devuelve tipo ConfigFunc
// entonces de esta manera podría leer la configuración
// del json.
func FromFile(path string) ConfigFunc {
	return func(opt *Client) {
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		decoder := json.NewDecoder(f)

		fop := struct {
			ConsultaIP string `json:"consulta_ip"`
			ConnString string `json:"conn_string"`
		}{}
		err = decoder.Decode(&fop)
		if err != nil {
			panic(err)
		}
		opt.consultaIP = fop.ConsultaIP
	}
}

// FromEnv lee la configuración
// de las variables de entorno
// y la combina con las existentes.
func FromEnv() ConfigFunc {
	return func(opt *Client) {
		connStr, exist := os.LookupEnv("CONN_DB")
		if exist {
			opt.connString = connStr
		}
	}
}

func NewClient(opts ...ConfigFunc) *Client {
	client := defaultClient
	for _, val := range opts {
		val(&client)
	}
	return &client
}

func main() {
	client := NewClient(FromFile("config.json"), FromEnv())
	fmt.Println(client.String())
}
