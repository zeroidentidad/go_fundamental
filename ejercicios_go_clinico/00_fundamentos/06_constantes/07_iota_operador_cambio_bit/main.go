package main

import "fmt"

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)  -> 2^10
	MB = 1 << (iota * 10) // 1 << (2 * 10)  -> 4^10
	GB = 1 << (iota * 10) // 1 << (3 * 10)  -> 8^10
	TB = 1 << (iota * 10) // 1 << (4 * 10)  -> 16^10
)

func main() {
	fmt.Println("binario\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}

/*
Operadores de cambio a nivel de bit . x << y significa x × 2y , mientras que x >> ysignifica x  × 2−y o, equivalentemente, x ÷ 2 y .
Los << y >> de Go son similares a los turnos (es decir, división o multiplicación por una potencia de 2)
*/
