package main

import "fmt"

func main() {

	miSlice := []string{"Lunes", "Martes"}
	miOtroSlice := []string{"Miercoles", "Jueves", "Viernes"}

	miSlice = append(miSlice, miOtroSlice...)

	fmt.Println(miSlice)
}
