package main

import (
	"fmt"
)

func main() {
	planetas := []string{
		"Mercurio",
		"Venus",
		"Tierra",
		"Marte",
		"Jupiter",
		"Saturno",
		"Urano",
		"Neptuno",
		"Pluton",
	}

	fmt.Println(cap(planetas[6:9]), len(planetas[6:9]), planetas[6:9])

	//newSlice := planetas[:]
	hyperspace(planetas)
	fmt.Println(planetas)

	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s = append(s, "d")

	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("cp:", c)
}

func hyperspace(planetas []string) {
	for i, _ := range planetas {
		planetas[i] = "Planeta " + planetas[i]
		//fmt.Println(i, planetas[i])
	}
}
