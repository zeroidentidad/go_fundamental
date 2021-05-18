package eightball

import (
	"math/rand"
	"time"
)

// randIntGenerator es una interfaz que incluye Intn,
// un método en el paquete math/rand incorporado.
// Esto permite simular el paquete math/rand en las pruebas.
type randIntGenerator interface {
	Intn(int) int
}

// EightBall simula una bola 8 mágica muy simple,
// un objeto mágico que predice el futuro respondiendo
// preguntas de sí/no.
type EightBall struct {
	rand randIntGenerator
}

// NewEightBall devuelve una nueva EightBall.
func New() *EightBall {
	return &EightBall{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Answer devuelve una respuesta mágica de eightball
// basada en un resultado aleatorio proporcionado por randomGenerator.
// Solo admite cuatro posibles respuestas.
func (e EightBall) Answer(s string) string {
	n := e.rand.Intn(3)
	switch n {
	case 0:
		return "Definitivamente no"
	case 1:

		return "Quizás"

	case 2:

		return "Si"

	default:

		return "Absolutamente"

	}

}
