package main

import "fmt"

func main() {
	m := map[string]int{
		"Batman": 32,
		"Robin":  27,
	}
	fmt.Println(m)

	fmt.Println(m["Batman"])

	fmt.Println(m["Eduar"])

	v, ok := m["Eduar"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Eduar"]; ok {
		fmt.Println("Impresión desde el if", v)
	}

}
