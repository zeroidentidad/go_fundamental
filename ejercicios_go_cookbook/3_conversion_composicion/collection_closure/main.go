package main

import (
	"fmt"

	"./collections"
)

func main() {
	ws := []collections.WorkWith{
		collections.WorkWith{"Ejemplo", 1},
		collections.WorkWith{"Ejemplo 2", 2},
	}

	fmt.Printf("Initial list: %#v\n", ws)

	// primera lista en minúsculas
	ws = collections.Map(ws, collections.LowerCaseData)
	fmt.Printf("Despues de LowerCaseData Map: %#v\n", ws)

	// incrementar todas las versiones
	ws = collections.Map(ws, collections.IncrementVersion)
	fmt.Printf("Despues de IncrementVersion Map: %#v\n", ws)

	// eliminar todas las versiones anteriores a la 3
	ws = collections.Filter(ws, collections.OldVersion(3))
	fmt.Printf("Despues de OldVersion Filter: %#v\n", ws)
}
