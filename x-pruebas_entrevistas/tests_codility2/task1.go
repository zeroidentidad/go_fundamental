package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := "The quick brown fox jumps over the lazy dog"

	fmt.Println(Solution(msg, 39))

	msg2 := "Codility We test coders"

	fmt.Println(Solution(msg2, 14))

	msg3 := "To crop or not to crop"

	fmt.Println(Solution(msg3, 21))
}

func Solution(message string, K int) string {
	values := strings.Fields(message)
	var result string
	var count = len(values) - 1

	for i, _ := range values {
		if count < K-1 {
			result += " " + values[i]
		} else {
			break
		}
		count += len(values[i])
	}

	return strings.TrimSpace(result)
}
