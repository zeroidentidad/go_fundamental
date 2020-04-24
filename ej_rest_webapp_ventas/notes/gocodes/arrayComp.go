package main

import (
	"fmt"
)

func main() {
	planetas := [...]string{
		"Mercurio",
		"Venus",
		"Tierra",
		"Marte",
		"Jupiter",
	}

	for i := 0; i < len(planetas); i++ {
		fmt.Println(i, planetas[i])
	}

	for i, planeta := range planetas {
		fmt.Println(i, planeta)
	}

	OverrideValues(planetas)
	fmt.Println(planetas)
}

func OverrideValues(planetas [5]string) {
	for i, _ := range planetas {
		planetas[i] = "New " + planetas[i]
		fmt.Println(i, planetas[i])
	}
}
