package main

type Saiyan struct {
	Name  string
	Power int
}

func main() {

	/*goku := &Saiyan{"Goku", 9001}
	goku.Super()
	fmt.Println(goku.Power)*/

	goku := new(Saiyan)
	// same as
	goku1 := &Saiyan{}

	goku2 := new(Saiyan)
	goku2.name = "goku"
	goku2.power = 9001
	//vs
	goku3 := &Saiyan{
		name:  "goku",
		power: 9000,
	}
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
