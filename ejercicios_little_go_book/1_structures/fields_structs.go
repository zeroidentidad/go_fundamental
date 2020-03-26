package main

import "fmt"

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

func main() {

	goku := &Saiyan{Name: "Goku"}
	goku.Power = 9000

	Super(goku)
	fmt.Println(goku.Power)

	gohan := &Saiyan{
		Name:  "Gohan",
		Power: 1000,
		Father: &Saiyan{
			Name:   "Goku",
			Power:  9001,
			Father: nil,
		},
	}
	fmt.Println(gohan.Father)
}

func Super(s *Saiyan) {
	s.Power += 10000
}
