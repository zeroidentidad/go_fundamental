package main

import "fmt"

func main() {
	m := map[string]int{
		"Batman": 32,
		"Robin":  27,
	}
	fmt.Println(m)

	delete(m, "Robin")
	fmt.Println(m)

	if v, ok := m["Robin"]; ok {
		fmt.Println("Se borró la llave con valor", v)
		delete(m, "Robin")
	}
	fmt.Println("No existe la llave especificada")
	fmt.Println(m)

}
