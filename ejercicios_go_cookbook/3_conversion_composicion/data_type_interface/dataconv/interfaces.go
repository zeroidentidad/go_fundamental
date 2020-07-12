package dataconv

import "fmt"

// CheckType se imprimirá en función del tipo de interfaz
func CheckType(s interface{}) {
	switch s.(type) {
	case string:
		fmt.Println("Es string!")
	case int:
		fmt.Println("Es int!")
	default:
		fmt.Println("No estoy seguro de que es...")
	}
}

// Interfaces demuestra la conversión de interfaces anónimas a tipos
func Interfaces() {
	CheckType("test")
	CheckType(1)
	CheckType(false)

	var i interface{}
	i = "test"

	// comprobar manualmente una interfaz
	if val, ok := i.(string); ok {
		fmt.Println("val es", val)
	}

	// este debería fallar
	if _, ok := i.(int); !ok {
		fmt.Println("oh! contento de haber controlado esto")
	}
}
