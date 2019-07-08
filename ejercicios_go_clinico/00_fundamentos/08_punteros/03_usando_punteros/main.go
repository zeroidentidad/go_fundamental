package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0xc0000180f0

	var b = &a
	fmt.Println(b)  // 0xc0000180f0
	fmt.Println(*b) // 43

	*b = 42        // b dice: "El valor en esta dirección, cámbielo a 42"
	fmt.Println(a) // 42

	// esto es útil
	// podemos pasar una dirección de memoria en lugar de un conjunto de valores (pasar una referencia)
	// y luego podemos cambiar el valor de lo que esté almacenado en esa dirección de memoria
	// Esto hace que nuestros programas sean de más redimiento.
	// No tenemos que pasar grandes cantidades de datos.
	// solo tenemos que pasar direcciones

	// todo es PASO POR VALOR en go, por cierto
	// cuando pasamos una dirección de memoria, estamos pasando un valor
}
