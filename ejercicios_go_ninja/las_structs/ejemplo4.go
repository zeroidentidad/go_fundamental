package main

import (
	"fmt"
)

type persona struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := persona{
		first: "Zero",
		last:  "Ferrer",
		age:   31,
	}

	fmt.Println(p1)
}
