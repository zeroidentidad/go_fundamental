package main

import "fmt"

func main() {
	fmt.Println(Calculo(5, 46))
	fmt.Println(Calculo(2, 23, 45, 68))
	fmt.Println(Calculo(5))
	fmt.Println(Calculo(5, 46, 17, 25, 26, 98, 17))

}

func uno(numero int) (z int) {
	z = numero * 2
	return
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func Calculo(numero ...int) int {
	total := 0
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item %d \n", item)
	}

	return total
}
