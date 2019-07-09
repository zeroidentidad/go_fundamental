package main

import "fmt"

func main() {

	miSlice := make([]int, 0, 3)

	fmt.Println("-----------------")
	fmt.Println(miSlice)
	fmt.Println(len(miSlice))
	fmt.Println(cap(miSlice))
	fmt.Println("-----------------")

	for i := 0; i < 80; i++ {
		miSlice = append(miSlice, i)
		fmt.Println("Longitud:", len(miSlice), "Capacidad:", cap(miSlice), "Valor: ", miSlice[i])
	}
}
