package errorwrap

import (
	"fmt"

	"github.com/pkg/errors"
)

// Unwrap desenvolverá un error
// y le escribirá una aserción
func Unwrap() {

	err := error(ErrorTyped{errors.New("ocurrió un error")})
	err = errors.Wrap(err, "wrapped")

	fmt.Println("wrapped error: ", err)

	// se puede manejar muchos tipos de error
	switch errors.Cause(err).(type) {
	case ErrorTyped:
		fmt.Println("Ocurrió un error tipado: ", err)
	default:
		fmt.Println("Un error desconocido ocurrió")
	}
}

// StackTrace imprimirá toda la pila del error
func StackTrace() {
	err := error(ErrorTyped{errors.New("ocurrió un error")})
	err = errors.Wrap(err, "wrapped")

	fmt.Printf("%+v\n", err)
}
