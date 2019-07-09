package main

import "fmt"

func main() {
	var x [58]string

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i) // 65 - 65 = string(0) = 'A'
	}

	fmt.Println(x)
	fmt.Println(x[42])
	fmt.Println(x[0])
}
