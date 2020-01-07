package main

import "fmt"

var tabla [10]int
var matriz [5][7]int

func main() {
	tabla[0] = 1
	tabla[5] = 15
	fmt.Println(tabla)

	matriz[3][5] = 1
	fmt.Println(matriz)

	// slices
	matrix := []int{2, 5, 4}
	fmt.Println(matrix)
	variante2()
	variante3()
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[2:4]

	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i <= 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d \n", len(nums), cap(nums))

}
