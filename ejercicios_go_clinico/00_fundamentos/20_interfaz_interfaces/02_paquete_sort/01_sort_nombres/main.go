package main

import (
	"fmt"
	"sort"
)

type personas []string

func (p personas) Len() int           { return len(p) }           // longitud
func (p personas) Swap(i, j int)      { p[i], p[j] = p[j], p[i] } // intercmabio
func (p personas) Less(i, j int) bool { return p[i] < p[j] }      // inferior

func main() {
	grupoEstudio := personas{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(grupoEstudio)
	sort.Sort(grupoEstudio)
	fmt.Println(grupoEstudio)

}

// https://golang.org/pkg/sort/#Sort
// https://golang.org/pkg/sort/#Interface
