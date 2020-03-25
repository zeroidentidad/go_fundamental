package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	/*jesus := Saiyan{
		Name:  "Zero",
		Power: 9000,
	}

	gohan := Saiyan{"Gohan", 9000}*/

	goku := &Saiyan{Name: "Goku"}
	goku.Power = 9000

	Super(goku)
	fmt.Println(goku.Power)
}

func Super(s *Saiyan) {
	//s.Power += 10000
	s = &Saiyan{"Gohan", 1000}
}
