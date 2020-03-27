package main

import (
	"fmt"
)

func main() {
	stra := "the spice must flow"
	byts := []byte(stra)
	strb := string(byts)

	fmt.Println(strb)

	fmt.Println(len("�"))
}
