package basicerrors

import (
	"errors"
	"fmt"
)

// ErrorValue es una forma de hacer un error de nivel
// de paquete para verificar. Es decir. if err == ErrorValue
var ErrorValue = errors.New("este es un error escrito")

// TypedError es una forma de cometer un tipo de error
// que puede hacer un err.(Type) == ErrorValue
type TypedError struct {
	error
}

//BasicErrors muestra algunas formas de crear errores
func BasicErrors() {
	err := errors.New("esta es una manera rápida y fácil de crear un error")
	fmt.Println("errors.New: ", err)

	err = fmt.Errorf("ocurrió un error: %s", "alguna cosa")
	fmt.Println("fmt.Errorf: ", err)

	err = ErrorValue
	fmt.Println("valor de error: ", err)

	err = TypedError{errors.New("error escrito")}
	fmt.Println("error escrito: ", err)

}
