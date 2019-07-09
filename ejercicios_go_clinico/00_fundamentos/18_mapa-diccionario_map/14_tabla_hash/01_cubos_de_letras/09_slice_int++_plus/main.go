package main

import "fmt"

func main() {
	cubos := make([]int, 1)
	fmt.Println(cubos[0])
	cubos[0] = 42
	fmt.Println(cubos[0])
	cubos[0]++
	fmt.Println(cubos[0])
}
