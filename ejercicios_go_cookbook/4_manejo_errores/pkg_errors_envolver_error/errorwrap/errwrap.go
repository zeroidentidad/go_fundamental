package errorwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

// WrappedError demuestra el ajuste de errores
// y la anotación de un error
func WrappedError(e error) error {
	return errors.Wrap(e, "Se produjo un error en WrappedError")
}

// ErrorTyped es un error que se puede verificar
type ErrorTyped struct {
	error
}

// Wrap shows what happens when we wrap an error
func Wrap() {
	e := errors.New("standard error")

	fmt.Println("Error regular - ", WrappedError(e))

	fmt.Println("Error tipado - ", WrappedError(ErrorTyped{errors.New("typed error")}))

	fmt.Println("Nil -", WrappedError(nil))

}
