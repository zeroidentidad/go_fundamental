package main

import "fmt"

func isSorted(A []int) bool {
	n := len(A)

	if n == 1 {
		return true
	}

	if A[n-1] < A[n-2] {
		return false
	}

	return isSorted(A[:n-1])
}

func main() {

	A := []int{10, 20, 23, 23, 45, 78, 88}
	fmt.Println(isSorted(A))

	A = []int{10, 20, 3, 23, 45, 78, 88}
	fmt.Println(isSorted(A))
}
