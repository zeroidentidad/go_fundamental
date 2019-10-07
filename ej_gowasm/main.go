package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	variable := "Golang"
	js.Global().Get("document").Get("body").Set("innerHTML", "Hola desde la web con "+variable)
	fmt.Println("Hola en la terminal")
}
