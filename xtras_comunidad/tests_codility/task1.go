package main

import "fmt"

func main() {
	coins := []int{1, 1, 1, 1, 1, 0, 1}

	fmt.Println(Solution(coins))
}

func Solution(A []int) int {
	// write your code in Go 1.4
	var alternates int

	for i := 0; i < len(A)-1; i++ {
		if A[i] == A[i+1] {
			alternates++
			A[i+1] = reverse(A[i])
		}
	}

	fmt.Println(A)

	return alternates
}

func reverse(item int) int {
	var value int
	if item == 0 {
		value = 1
	} else if item == 1 {
		value = 0
	}

	return value
}
