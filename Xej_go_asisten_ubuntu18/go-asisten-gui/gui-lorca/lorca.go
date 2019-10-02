package gui

import (
	"github.com/zserge/lorca"
)

// NewLorcaHTTP lanzar ventana lorca http
func NewLorcaHTTP(route string, width int, height int) (lorca.UI, error) {
	return lorca.New(route, "", width, height)
}
