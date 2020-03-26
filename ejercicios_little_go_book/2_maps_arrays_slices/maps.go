package main

import (
	"fmt"
)

func main() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	fmt.Println(lookup)

	power, exists := lookup["vegeta"]
	fmt.Println(power, exists)

	total := len(lookup)
	fmt.Println(total)

	delete(lookup, "goku")
	fmt.Println(lookup)

	// grow dynamically
	lookup2 := make(map[string]int, 100)
	fmt.Println(lookup2)

	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}

	goku2 := Saiyan2{
		Name:    "Goku2",
		Friends: make(map[string]Saiyan2),
	}

	goku.Friends["krillin"] = &Saiyan{"Krillin", nil}
	fmt.Println(goku.Name)
	for _, value := range goku.Friends {
		fmt.Println(value)
	}

	goku2.Friends["krillin2"] = Saiyan2{"Krillin2", nil}
	fmt.Println(goku2.Name)
	for _, value := range goku2.Friends {
		fmt.Println(value)
	}

	lookup3 := map[string]int{
		"goku":  9001,
		"gohan": 2044,
	}

	for key, value := range lookup3 {
		fmt.Println(key, value)
	}
}

// as field of struct
type Saiyan struct {
	Name    string
	Friends map[string]*Saiyan
}

type Saiyan2 struct {
	Name    string
	Friends map[string]Saiyan2
}
