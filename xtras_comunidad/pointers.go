package main

import "fmt"

func main() {
	var a int = 10
	var p *int = &a // p de tipo *int apunta a, a

	fmt.Println(p)  // "0x0001"
	fmt.Println(*p) // "10"

	*p = 20        // actualizar indirectamente a
	fmt.Println(a) // "20"
}

func p() {
	var n *int
	var x, y int

	fmt.Println(n)         // "<nil>"
	fmt.Println(n == nil)  // "true" (n es nil)
	fmt.Println(x == y)    // "true" (x e y son ambos cero)
	fmt.Println(&x == &x)  // "true" (*x es igual a sí mismo)
	fmt.Println(&x == &y)  // "false" (diferentes variables)
	fmt.Println(&x == nil) // "false" (*x no es nil)
}
