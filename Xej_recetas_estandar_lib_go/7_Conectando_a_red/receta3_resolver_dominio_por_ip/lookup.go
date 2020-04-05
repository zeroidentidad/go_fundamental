package main

import (
	"fmt"
	"net"
)

func main() {

	// Resolver por IP
	addrs, err := net.LookupAddr("127.0.0.1")
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}

	//Resolver por direccion
	ips, err := net.LookupIP("google.com")
	if err != nil {
		panic(err)
	}

	for _, ip := range ips {
		fmt.Println(ip.String())
	}
}
