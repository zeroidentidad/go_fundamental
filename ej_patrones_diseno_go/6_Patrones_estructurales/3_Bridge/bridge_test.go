package bridge

import (
	"fmt"
	"testing"
)

// Test de uso del patrón
func TestPattern(t *testing.T) {
	telefonoA := &Telefono{"0115161", &Dispositivo{}}
	telefonoA.SetConexion(&Red4G{})
	fmt.Printf("%s\n", telefonoA.ConectarInternet())

	telefonoB := &Telefono{"0117854", &Dispositivo{}}
	telefonoB.SetConexion(&RedWiFi{})
	fmt.Printf("%s\n", telefonoB.ConectarInternet())

	tablet := &Tablet{&Dispositivo{}}
	tablet.SetConexion(&RedWiFi{})
	fmt.Printf("%s\n", tablet.ConectarInternet())
}
