package visitor

import (
	"fmt"
	"testing"
)

func TestPattern(t *testing.T) {
	// desde visitar
	elementoArma0 := &ElementoArma{}
	elementoSuperpoder0 := &ElementoSuperpoder{}

	visitanteNivel0 := &VisitanteNivel0{}
	visitanteNivel0.VisitarArma(elementoArma0)
	visitanteNivel0.VisitarSuperpoder(elementoSuperpoder0)

	fmt.Printf("El visitante Nivel 0 tiene la siguiente arma de batalla: %s\n", elementoArma0.Arma)
	fmt.Printf("El visitante Nivel 0 tiene el siguiente superpoder: %s\n", elementoSuperpoder0.Poder)

	elementoArma1 := &ElementoArma{}
	elementoSuperpoder1 := &ElementoSuperpoder{}

	visitanteNivel1 := &VisitanteNivel1{}
	visitanteNivel1.VisitarArma(elementoArma1)
	visitanteNivel1.VisitarSuperpoder(elementoSuperpoder1)

	fmt.Printf("El visitante Nivel 1 tiene la siguiente arma de batalla: %s\n", elementoArma1.Arma)
	fmt.Printf("El visitante Nivel 1 tiene el siguiente superpoder: %s\n", elementoSuperpoder1.Poder)

	// desde aceptar
	visitanteNivel0 = &VisitanteNivel0{}
	elementoArma0 = &ElementoArma{}
	elementoArma0.Aceptar(visitanteNivel0)

	fmt.Printf("El elemento Arma aceptada por un visitante Nivel 0 es: %s\n", elementoArma0.Arma)
}
