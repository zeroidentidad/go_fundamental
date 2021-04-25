package main

import (
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {

	message.Set(language.Spanish, "%d elementos para hacer",
		plural.Selectf(1, "%d",
			"=0", "sin elementos para hacer",
			plural.One, "un elemento para hacer",
			"<100", "%[1]d elementos para hacer",
			plural.Other, "muchos elementos para hacer",
		))

	message.Set(language.Spanish, "El promedio es %.2f",
		plural.Selectf(1, "%.2f",
			"<1", "El promedio es cero",
			"=1", "El promedio es uno",
			plural.Other, "El promedio es %[1]f ",
		))

	prt := message.NewPrinter(language.Spanish)
	prt.Printf("%d elementos para hacer", 0)
	prt.Println()
	prt.Printf("%d elementos para hacer", 1)
	prt.Println()
	prt.Printf("%d elementos para hacer", 10)
	prt.Println()
	prt.Printf("%d elementos para hacer", 1000)
	prt.Println()

	prt.Printf("El promedio es %.2f", 0.8)
	prt.Println()
	prt.Printf("El promedio es %.2f", 1.0)
	prt.Println()
	prt.Printf("El promedio es %.2f", 10.0)
	prt.Println()

}
