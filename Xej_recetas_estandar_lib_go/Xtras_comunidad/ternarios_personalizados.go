package main

import (
	"fmt"
)

func main() {
	a := ternaryInt64(5 > 88, 13, 0)
	b := ternaryString(3 > 1, "correcto", "incorrecto")
	fmt.Println(a)
	fmt.Println(b)
}

func ternaryInt64(a bool, b, c int64) int64 {
	if a {
		return b
	}

	return c
}

func ternaryString(a bool, b, c string) string {
	if a {
		return b
	}

	return c
}

//package ternary

// Ternary te da una expresión ternaria simple pasando tipos de datos diferentes

func ternary(a bool, b, c interface{}) interface{} {
	if a {
		return b
	}

	return c
}
