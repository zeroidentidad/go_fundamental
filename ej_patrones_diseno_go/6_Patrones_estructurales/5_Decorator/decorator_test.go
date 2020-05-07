package decorator

import (
	"fmt"
	"testing"
)

// Test de uso del patrón
func TestPattern(t *testing.T) {
	cafe := &Cafe{}
	fmt.Printf("Detalle: %s - Importe $%d\n", cafe.GetDetalle(), cafe.GetCosto())

	cafeConCrema := &CafeConCrema{&CafeDecorador{cafe}}
	fmt.Printf("Detalle: %s - Importe $%d\n", cafeConCrema.GetDetalle(), cafeConCrema.GetCosto())

	cafeConCremaConCanela := &CafeConCanela{&CafeDecorador{cafeConCrema}}
	fmt.Printf("Detalle: %s - Importe $%d\n", cafeConCremaConCanela.GetDetalle(), cafeConCremaConCanela.GetCosto())
}
