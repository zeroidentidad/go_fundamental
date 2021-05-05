package main

import "fmt"

func main() {
	nums := []int{1, 2, 6, 7}

	fmt.Println(Solution(nums))
}

func Solution(A []int) int {
	// write your code in Go 1.4
	var n int = A[0]
	for index, value := range A {
		if A[index] > n {
			n = value
		}
	}

	return n + 1
}
