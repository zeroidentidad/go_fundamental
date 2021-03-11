package main

import (
	"fmt"
)

func main() {
	const height uint8 = 15
	header(height)
	base(height)
}

func header(height uint8) {
	for i := 0; i <= int(height); i++ {
		for j := int(height); j > i; j-- {
			fmt.Print(" ")
		}
		for k := 0; k < i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func base(height uint8) {
	for i := 0; i <= int(height)/5; i++ {
		for j := int(height) / 3; j > 0; j-- {
			fmt.Print("  ")
		}
		for k := 0; k < int(height)/3; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
