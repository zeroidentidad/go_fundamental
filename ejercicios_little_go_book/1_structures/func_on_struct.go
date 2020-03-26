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

	goku := &Saiyan{"Goku", 9001}
	goku.Super()            // recibe struct
	fmt.Println(goku.Power) // will print 19001
}

func (s *Saiyan) Super() {
	s.Power += 10000
}
