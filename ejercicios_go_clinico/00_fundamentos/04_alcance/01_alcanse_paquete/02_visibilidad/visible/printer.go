package visible

import "fmt"

// PrintVar se exporta porque comienza con una letra mayúscula.
func PrintVar() {
	fmt.Println(MiNombre)
	fmt.Println(tuNombre)
}
