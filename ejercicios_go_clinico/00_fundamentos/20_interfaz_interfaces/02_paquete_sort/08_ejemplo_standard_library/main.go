package main

import (
	"fmt"
	"sort"
)

type persona struct {
	Nombre string
	Edad   int
}

func (p persona) String() string {
	return fmt.Sprintf("YAYAYA %s: %d", p.Nombre, p.Edad)
}

// PorEdad implements sort.Interface para []persona basaso en
// en el campo Edad.
type PorEdad []persona

func (a PorEdad) Len() int           { return len(a) }                // longitud
func (a PorEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }      // intercambio
func (a PorEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad } //inferior

//func (a PorEdad) Less(i, j int) bool { return a[i].Nombre < a[j].Nombre }

func main() {
	personas := []persona{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(personas[0])
	fmt.Println(personas)
	sort.Sort(PorEdad(personas))
	fmt.Println(personas)

}

// https://golang.org/pkg/sort/#Sort
// https://golang.org/pkg/sort/#Interface

// String() string
// https://golang.org/doc/effective_go.html#printing
