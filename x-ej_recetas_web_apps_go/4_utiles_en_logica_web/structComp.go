package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address   string
}

func main() {

	var persona Person
	persona.FirstName = "Jesus"
	persona.LastName = "Ferrer"
	persona.Age = 27
	persona.Address = "Mexico"

	persona2 := Person{
		FirstName: "Jesus",
		LastName:  "Ferrer",
		Age:       27,
		Address:   "Mexico",
	}

	persona3 := Person{"Jesus", "Ferrer", 27, "Mexico"}

	persona4 := persona3
	persona4.FirstName = "Liliam"
	persona4.Age = 0

	personas := []Person{
		{"Jesus1", "Ferrer", 27, "Mexico"},
		{"Jesus2", "Ferrer", 28, "Mexico"},
		{"Jesus3", "Ferrer", 29, "Mexico"},
	}

	fmt.Println(persona)
	fmt.Println(persona2)
	fmt.Println(persona3)
	fmt.Printf("%+v\n", persona4)
	fmt.Printf("%+v\n", personas)
}
