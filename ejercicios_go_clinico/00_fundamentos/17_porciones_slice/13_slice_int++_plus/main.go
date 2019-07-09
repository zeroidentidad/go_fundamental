package main

import "fmt"

func main() {
	miSlice := make([]int, 1)
	fmt.Println(miSlice[0])
	miSlice[0] = 7
	fmt.Println(miSlice[0])
	miSlice[0]++
	fmt.Println(miSlice[0])
}
