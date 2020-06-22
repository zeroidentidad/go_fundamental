package main

import (
	"flag"
	"fmt"

	"../flags"
)

func main() {
	// inicializa nuestra configuración
	c := flags.Config{}
	c.Setup()

	// generalmente llamamos a esto desde main
	flag.Parse()

	fmt.Println(c.GetMessage())
}
