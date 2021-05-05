package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Solution(3, 7))
}

func Solution(A int, B int) int {
	// write your code in Go 1.4
	value := strconv.Itoa(A * B)
	toBin, _ := convert(value, 10, 2)

	fmt.Println(toBin)

	var count int
	bit := []byte("1")
	for i := 0; i < len(toBin); i++ {
		if toBin[i] == bit[0] {
			count++
		}
	}

	return count
}

func convert(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}
