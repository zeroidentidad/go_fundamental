package main

import (
	"flag"
	"fmt"
)

func main() {
	port := flag.Int("p", 8000,
		"specify port. default is 8000.")

	enable := flag.Bool("e", false,
		"specify enable. default is false.")

	name := flag.String("n", "blank",
		"specify name. default is blank.")

	flag.Parse()

	fmt.Println("port = ", *port)
	fmt.Println("enable = ", *enable)
	fmt.Println("name = ", *name)
}
