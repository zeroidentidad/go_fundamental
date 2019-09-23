package main

import (
	"fmt"
)

func main() {
	var asignaturas map[string]string
	fmt.Println(asignaturas)
	deportes := map[string]int{
		"Tiro con arco": 20,
	}
	deportes["Tiro con espada"] = 40
	fmt.Println(deportes)
	resultado, ok := deportes["Tiro con arco"]
	fmt.Println(resultado, ok)
	delete(deportes, "Tiro con javalina")
	fmt.Println(deportes)
	for clave, valor := range deportes {
		fmt.Println(clave, valor)
	}
}
