package main

import "fmt"

func main() {

	miSaludo := map[string]string{
		"cero": "Good morning!",
		"uno":  "Bonjour!",
		"dos":  "Buenos dias!",
		"tres": "Bongiorno!",
	}

	fmt.Println(miSaludo)
	delete(miSaludo, "dos")
	fmt.Println(miSaludo)
}
