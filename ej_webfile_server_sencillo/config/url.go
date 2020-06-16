package config

import (
	"fmt"
	"net"
)

func UrlExample() {
	fmt.Println("------------------------------------------")
	fmt.Println("URL de ejemplo para acceder:")
	fmt.Println("------------------------------------------")
	fmt.Println("	http://" + LocalIP() + ":" + PortHttp(false))
}

func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
