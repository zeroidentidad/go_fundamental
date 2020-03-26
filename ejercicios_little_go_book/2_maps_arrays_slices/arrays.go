package main

import (
	"fmt"
)

func main() {
	//var scores [10] int
	//scores[0]= 399

	scores := [4]int{9001, 9333, 212, 33}

	fmt.Println("len", len(scores))

	for index, value := range scores {
		fmt.Println(index, ":", value)
	}
}
