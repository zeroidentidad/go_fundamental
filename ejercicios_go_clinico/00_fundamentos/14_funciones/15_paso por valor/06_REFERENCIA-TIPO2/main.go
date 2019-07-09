package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ]
	cambiaMe(m)
	fmt.Println(m) // [Jesus]
}

func cambiaMe(z []string) {
	z[0] = "Jesus"
	fmt.Println(z) // [Jesus]
}
