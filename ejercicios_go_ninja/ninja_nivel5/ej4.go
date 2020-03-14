package main

import (
	"fmt"
)

func main() {

	s := struct {
		nombre     string
		amigos     map[string]int
		bebidasFav []string
	}{
		nombre: "Zero",
		amigos: map[string]int{
			"Carito":  222,
			"Adriana": 444,
			"Condor":  666,
		},
		bebidasFav: []string{
			"agua",
			"chicha",
		},
	}
	fmt.Println(s.nombre)
	fmt.Println(s.amigos)
	fmt.Println(s.bebidasFav)

	for k, v := range s.amigos {
		fmt.Println(k, v)
	}

	for i, val := range s.bebidasFav {
		fmt.Println(i, val)
	}
}
