package main

import (
	"fmt"
)

func main() {
	temperature := map[string]int{
		"Tierra": 15,
		"Marte":  -65,
	}

	temperature["Tierra"] = 50

	fmt.Println(temperature["Marte"])
	fmt.Println(temperature["Tierra"])

	temperature2 := temperature
	temperature2["Tierra"] = 1000
	fmt.Println(temperature)

	m := make(map[string]int)
	m["k1"] = 17
	m["k2"] = 112
	m["k3"] = 7
	m["k4"] = 45

	delete(m, "k2")
	fmt.Println(m)

	for _, val := range m {
		fmt.Println(val)
	}
}
