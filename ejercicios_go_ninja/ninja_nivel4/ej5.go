package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	x = append(x[:3], x[6:]...) // [42, 43, 44, 48, 49, 50, 51]
	fmt.Println(x)
}
