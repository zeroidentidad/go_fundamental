package main

import "fmt"

func main() {

	miSlice := []string{"Lunes", "Martes"}
	miOtroSlice := []string{"Jueves", "Viernes", "Sabado"}

	miSlice = append(miSlice, miOtroSlice...)
	fmt.Println(miSlice)

	miSlice = append(miSlice[:2], miSlice[3:]...)
	fmt.Println(miSlice)

}
