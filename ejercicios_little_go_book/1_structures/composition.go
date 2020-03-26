package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type Saiyan struct {
	*Person
	Power int
}

func main() {
	// to use it
	goku := &Saiyan{
		Person: &Person{"Goku"},
		Power:  9001,
	}
	goku.Person.Introduce()
	goku.Introduce()

	fmt.Println(goku.Name)
	fmt.Println(goku.Person.Name)
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

// overloading
func (s *Saiyan) Introduce() {
	fmt.Printf("Hi, I'm %s. Ya!\n", s.Name)
}
