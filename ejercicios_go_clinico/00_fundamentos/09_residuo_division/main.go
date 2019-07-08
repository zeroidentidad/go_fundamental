package main

import "fmt"

func main() {
	x := 13 % 3
	fmt.Println(x)
	if x == 1 {
		fmt.Println("Impar")
	} else {
		fmt.Println("Par")
	}
	fmt.Println("----------------------------------")
	for i := 1; i <= 70; i++ {
		if i%2 == 1 {
			fmt.Println(i, "Impar")
		} else {
			fmt.Println(i, "Par")
		}
	}
}
