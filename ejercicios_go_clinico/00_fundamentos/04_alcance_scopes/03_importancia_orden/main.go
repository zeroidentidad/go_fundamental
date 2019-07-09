package main

import "fmt"

func main() {
	fmt.Println(x) //error
	fmt.Println(y)
	x := 42
}

var y = 42
