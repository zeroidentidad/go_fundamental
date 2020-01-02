package main

import (
	"net/http"

	"gopkg.in/unrolled/render.v1"
)

// Action, define una firma de función estándar para que la usemos al crear
// acciones del controlador. Una acción de controlador es básicamente solo un método asociado a
// un controlador.
type Action func(rw http.ResponseWriter, r *http.Request) error

// Este es nuestro Base Controller
type AppController struct{}

// La función de acción ayuda con el manejo de errores en un controlador
func (c *AppController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := a(rw, r); err != nil {
			http.Error(rw, err.Error(), 500)
		}
	})
}

type MyController struct {
	AppController
	*render.Render
}

func (c *MyController) Index(rw http.ResponseWriter, r *http.Request) error {
	c.JSON(rw, 200, map[string]string{"Hello": "JSON"})
	return nil
}

func main() {
	c := &MyController{Render: render.New(render.Options{})}
	http.ListenAndServe(":8080", c.Action(c.Index))
}
