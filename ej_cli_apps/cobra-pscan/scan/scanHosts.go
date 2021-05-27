// Package scan proporciona tipos y funciones para realizar
// escaneos de puertos TCP en una lista de hosts
package scan

import (
	"fmt"
	"net"
	"time"
)

// PortState representa el estado de un solo puerto TCP
type PortState struct {
	Port int
	Open state
}

type state bool

// String convierte valor booleano del estado en string legible
func (s state) String() string {
	if s {
		return "open"
	}

	return "closed"
}

// scanPort realiza un escaneo de puertos en un solo puerto TCP
func scanPort(host string, port int) PortState {
	p := PortState{
		Port: port,
	}

	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	scanConn, err := net.DialTimeout("tcp", address, 1*time.Second)

	if err != nil {
		return p
	}

	scanConn.Close()
	p.Open = true
	return p
}

// Results representa los resultados del análisis para un solo host
type Results struct {
	Host       string
	NotFound   bool
	PortStates []PortState
}

// Run realiza un escaneo de puertos en la lista de hosts
func Run(hl *HostsList, ports []int) []Results {
	res := make([]Results, 0, len(hl.Hosts))
	for _, h := range hl.Hosts {
		r := Results{
			Host: h,
		}

		if _, err := net.LookupHost(h); err != nil {
			r.NotFound = true
			res = append(res, r)
			continue
		}

		for _, p := range ports {
			r.PortStates = append(r.PortStates, scanPort(h, p))
		}

		res = append(res, r)
	}

	return res
}
