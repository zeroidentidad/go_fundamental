package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {

	goku := &Saiyan{"Goku", 9001}
	goku.Super()
	fmt.Println(goku.Power)
}

func (s *Saiyan) Super() {
	s.Power += 10000
}

func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}

func NewSaiyan2(name string, power int) Saiyan { // valid without return pointer
	return Saiyan{
		Name:  name,
		Power: power,
	}
}
