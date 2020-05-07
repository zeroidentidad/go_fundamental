package flyweight

import (
	"fmt"
	"testing"
)

// Test de uso del patrón
func TestPattern(t *testing.T) {
	pantalla := &PantallaFabrica{}

	fmt.Printf("%s\n", pantalla.ObtenerControl("BOTON", "BTN1", "100", "300"))
	fmt.Printf("%s\n", pantalla.ObtenerControl("BOTON", "BTN2", "200", "300"))
	fmt.Printf("%s\n", pantalla.ObtenerControl("BOTON", "BTN3", "300", "300"))
	fmt.Printf("%s\n", pantalla.ObtenerControl("PASSWORD", "PWD1", "500", "300"))
	fmt.Printf("%s\n", pantalla.ObtenerControl("BOTON", "BTN1", "400", "300"))
}
