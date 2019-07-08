package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	foo()
}

func foo() {
	// sin acceso a x
	// esto no compila
	fmt.Println(x)
}
