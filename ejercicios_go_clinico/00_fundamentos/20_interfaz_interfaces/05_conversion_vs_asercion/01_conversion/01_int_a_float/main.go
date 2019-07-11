package main

import "fmt"

func main() {
	var x = 12
	var y = 12.1230123
	fmt.Println(float64(x))
	fmt.Println(y + float64(x))
	// conversion: int to float64
}
