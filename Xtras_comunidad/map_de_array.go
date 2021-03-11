package main

import "fmt"

func main() {

	datos := make(map[string][]int)

	datos["key1"] = []int{10, 9, 8, 6, 7}
	datos["key2"] = []int{11, 12, 5, 4, 2}
	fmt.Println(datos)

	datos["key1"][0] = 100 // ok
	fmt.Println(datos)

	delete(datos, "key2") // no ok: datos["key2"][1]
	fmt.Println(datos)

	// para eliminar o agregar necesario re-slicing
	// de array independiente del mapa (más pasos)
}
