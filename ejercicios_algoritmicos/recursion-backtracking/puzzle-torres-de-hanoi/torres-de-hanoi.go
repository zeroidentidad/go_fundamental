package main

import "fmt"

func TOHv1(n int, from, to, temp string) {

	/* Si solo es 1 disco, hacer movimiento y regresar */
	if n == 1 {
		fmt.Println("Mover disco ", n, " de varilla ", from, " a varilla ", to)
		return
	}

	/* Mover los n-1 discos superiores de A a B, usando C como auxiliar */
	TOHv1(n-1, from, temp, to)

	/* Mover los discos restantes de A a C */
	fmt.Println("Mover disco ", n, " de varilla ", from, " a varilla ", to)

	/* Mover n-1 discos de B a C usando A como auxiliar */
	TOHv1(n-1, temp, to, from)
}

func TowersOfHanoi(n int) {
	TOHv1(n, "A", "C", "B")
}

func main() {
	TowersOfHanoi(3)
}
