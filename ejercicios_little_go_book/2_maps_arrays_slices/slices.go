package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	// formas de crear/declarar una slice
	scores := []int{1, 4, 293, 4, 9}
	fmt.Println(scores)

	scores2 := make([]int, 10)
	fmt.Println(scores2)

	scores3 := make([]int, 0, 10)
	scores3 = append(scores3, 5)
	fmt.Println(scores3)

	scores4 := make([]int, 0, 10)
	scores4 = scores4[0:8]
	scores4[7] = 9033
	fmt.Println(scores4)

	slices()

	slices2()

	commonways()

	goku := []*Saiyans{{Name: "goku", Power: 9000}, {Name: "zero", Power: 1000}}

	powers := extractPowers(goku)
	fmt.Println(powers)

	scores5 := []int{1, 2, 3, 4, 5}
	slice := scores5[2:4]
	slice[0] = 999
	fmt.Println(scores5)

	haystack := "the spice must flow"
	fmt.Println(strings.Index(haystack[5:], " ")) // start "p" index before "m"

	scores6 := []int{1, 2, 3, 4, 5}
	scores6 = scores6[:len(scores6)-1] // values of a slice except the last (Go no support negative values)
	fmt.Println(scores6)

	// remove a value from an unsorted slice
	scores7 := []int{1, 2, 3, 4, 5}
	scores7 = removeAtIndex(scores7, 2)
	fmt.Println(scores7)

	// copy use
	scores8 := make([]int, 100)

	for i := 0; i < 100; i++ {
		scores8[i] = int(rand.Int31n(1000))
	}

	sort.Ints(scores8)
	worst := make([]int, 5)
	copy(worst, scores8[:5])
	fmt.Println(worst)

}

func slices() {
	scores := make([]int, 0, 5)
	c := cap(scores)
	fmt.Println(c)
	for i := 0; i < 25; i++ {
		scores = append(scores, i)
		// si la capacidad ha cambiado,
		// Go hace crecer la matriz para acomodar nuevos datos
		if cap(scores) != c {
			c = cap(scores)
			fmt.Println(c)
		}
	}
}

func slices2() {
	scores := make([]int, 5)
	scores = append(scores, 9332)
	fmt.Println(scores)
}

func commonways() {
	names := []string{"leto", "zero", "paul"}
	checks := make([]bool, 5)
	var names2 []string
	scores := make([]int, 0, 10)
	fmt.Println(names, checks, names2, scores)
}

type Saiyans struct {
	Name  string
	Power int
}

func extractPowers(saiyans []*Saiyans) []int {
	powers := make([]int, len(saiyans))
	for index, saiyan := range saiyans {
		powers[index] = saiyan.Power
	}
	return powers
}

func extractPowers2(saiyans []*Saiyans) []int {
	powers := make([]int, 0, len(saiyans))
	for _, saiyan := range saiyans {
		powers = append(powers, saiyan.Power)
	}
	return powers
}

func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	//swap the last value and the value we want to remove
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}
