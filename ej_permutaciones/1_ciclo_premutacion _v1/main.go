package main

import "fmt"

func main() {
	options := []string{"Hola", "Yo", "Gente mala"}
	for _, x := range options {
		for _, y := range options {
			for _, z := range options {
				fmt.Println(x, y, z)
			}
		}
	}
}
