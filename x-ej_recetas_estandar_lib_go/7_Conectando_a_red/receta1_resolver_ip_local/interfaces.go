package main

import (
	"fmt"
	"net"
)

func main() {

	// Obtener todas las interfaces de red
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, interf := range interfaces {
		// Resolver direcciones para cada interfaz
		addrs, err := interf.Addrs()
		if err != nil {
			panic(err)
		}
		fmt.Println(interf.Name)
		for _, add := range addrs {
			if ip, ok := add.(*net.IPNet); ok {
				fmt.Printf("\t%v\n", ip)
			}
		}

	}

}
