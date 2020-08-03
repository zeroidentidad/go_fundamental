package basicerrors

import (
	"fmt"
)

// CustomError es una estructura que implementará la interfaz Error()
type CustomError struct {
	Result string
}

// implementar y devolver resultado
func (c CustomError) Error() string {
	return fmt.Sprintf("ocurrio un error; %s fue el resultado", c.Result)
}

// SomeFunc devuelve un error
func SomeFunc() error {
	c := CustomError{Result: "Esta Onda"}
	return c
}
