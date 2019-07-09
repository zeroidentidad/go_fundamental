package main

import "fmt"

type cliente struct {
	nombre string
	edad   int
}

func main() {
	c1 := cliente{"Jesus", 26}
	fmt.Println(&c1.nombre) // 0xc000090020

	cambiaMe(&c1)

	fmt.Println(c1)         // {Rocky 26}
	fmt.Println(&c1.nombre) // 0xc000090020
}

func cambiaMe(z *cliente) {
	fmt.Println(z)         // &{Jesus 26}
	fmt.Println(&z.nombre) // 0xc000090020
	z.nombre = "Rocky"
	fmt.Println(z)         // &{Rocky 26}
	fmt.Println(&z.nombre) // 0xc000090020

}
