package singleton

import (
	"fmt"
	"testing"
	"time"
)

// Test de uso del patrón
func TestPattern(t *testing.T) {
	fmt.Println("Todas las instancias Singleton tienen que tener el mismo número")

	fmt.Printf("Instancia Singleton: %d\n", GetInstancia().Tiempo)

	time.Sleep(1 * time.Second)

	fmt.Printf("Instancia Singleton: %d\n", GetInstancia().Tiempo)

	canalEspera := make(chan int64)

	go func() {
		time.Sleep(1 * time.Second)

		canalEspera <- GetInstancia().Tiempo
	}()

	fmt.Printf("Instancia Singleton: %d\n", <-canalEspera)
}
