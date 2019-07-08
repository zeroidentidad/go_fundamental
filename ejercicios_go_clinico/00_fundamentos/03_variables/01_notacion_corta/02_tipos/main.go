package main

import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hola"
	f := `Te gusta mi sombrero?`
	g := 'M'

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)
}
