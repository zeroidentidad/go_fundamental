package main

import "fmt"

func main() {

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	fmt.Println(food) // fuera del scope o alcanse

}