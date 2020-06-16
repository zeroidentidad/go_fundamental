package config

import (
	"fmt"
	"net"
)

func GetIPs() {
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	fmt.Println("> Host: " + GetHostame())
	fmt.Println("------------------------------------------")
	fmt.Println("IPs LAN de localhost:")
	fmt.Println("------------------------------------------")

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			panic(err)
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				if !v.IP.IsLoopback() && v.IP.To4() != nil {
					ip = v.IP
					fmt.Println(ip)
				}
			}
		}

	}
}
