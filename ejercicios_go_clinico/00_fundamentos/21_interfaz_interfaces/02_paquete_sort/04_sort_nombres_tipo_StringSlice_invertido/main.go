package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)

	// para experimentación para entender lo que está pasando:
	// descomenta y experimenta con el siguiente código:

	//	sort.Sort(sort.StringSlice(s))
	//	fmt.Println(s)
	//
	//	fmt.Printf("sólo s: %T\n", s)
	//	s = sort.StringSlice(s)
	//	fmt.Printf("sólo s: %T\n", s)
	//	t := sort.StringSlice(s)
	//	fmt.Printf("sólo t: %T\n", t)
	//
	//	fmt.Printf("s convertido a StringSlice: %T\n", sort.StringSlice(s))
	////	fmt.Printf("s ordenado: %T\n", sort.Sort(sort.StringSlice(s)))
	//	fmt.Printf("s invertido: %T\n", sort.Reverse(sort.StringSlice(s)))
	//	i := sort.Reverse(sort.StringSlice(s))
	//	fmt.Println(i)
	//	fmt.Printf("%T\n", i)
	//	sort.Sort(i)
	//	fmt.Println(s)
}

// https://golang.org/pkg/sort/#Sort
