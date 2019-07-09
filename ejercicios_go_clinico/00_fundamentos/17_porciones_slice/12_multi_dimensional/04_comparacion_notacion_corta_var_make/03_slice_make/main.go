package main

import (
	"fmt"
)

func main() {
	estudiante := make([]string, 35)
	estudiantes := make([][]string, 35)
	estudiante[0] = "Jesus"
	// estudiante = append(estudiante, "Jesus")
	fmt.Println(estudiante)
	fmt.Println(estudiantes)
}
