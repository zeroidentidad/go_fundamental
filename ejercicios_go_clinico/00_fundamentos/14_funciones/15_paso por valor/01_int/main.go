package main

import "fmt"

func main() {
	edad := 44
	cambiaMe(edad)
	fmt.Println(edad) // 44
}

func cambiaMe(z int) {
	fmt.Println(z)
	z = 24
}

// cuando se llama cambiaMe en la línea 8
// el valor 44 se pasa como un argumento
