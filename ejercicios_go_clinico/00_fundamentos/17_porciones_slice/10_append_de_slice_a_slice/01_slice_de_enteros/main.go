package main

import "fmt"

func main() {

	miSlice := []int{1, 2, 3, 4, 5}
	miOtroSlice := []int{6, 7, 8, 9}

	miSlice = append(miSlice, miOtroSlice...)

	fmt.Println(miSlice)
}
