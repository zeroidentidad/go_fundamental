package dataconv

import (
	"fmt"
	"strconv"
)

// Strconv demuestra algunas funciones strconv
func Strconv() error {
	//strconv es buena manera de convertir a... y, desde cadenas
	s := "1234"
	// podemos especificar la base (10) y la precisión de 64 bits
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	fmt.Println(res)

	// intentar hexadecimal
	res, err = strconv.ParseInt("FF", 16, 64)
	if err != nil {
		return err
	}

	fmt.Println(res)

	// Podemos hacer otras cosas útiles como:
	val, err := strconv.ParseBool("true")
	if err != nil {
		return err
	}

	fmt.Println(val)

	return nil
}
