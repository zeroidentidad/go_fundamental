package main

import (
	"fmt"
)

type persona struct {
	first string
	last  string
	age   int
}

type agenteSecreto struct {
	persona
	lpm bool
}

func main() {

	as1 := agenteSecreto{
		persona: persona{
			first: "James",
			last:  "Bond",
			age:   31,
		},
		lpm: true,
	}

	p2 := persona{
		first: "Condor",
		last:  "Pérez",
		age:   55,
	}

	fmt.Println(as1)
	fmt.Println(p2)
	fmt.Println(as1.first, as1.last, as1.age)
	fmt.Println(p2.first, p2.last, p2.age)

}
