package eightball

import (
	"testing"
)

type fixedRandIntGenerator struct {
	// tnúmero que debería generarse "aleatoriamente"
	randomNum int
	// registrar el parámetro con el que se llama a Intn
	calledWithN int
}

func (g *fixedRandIntGenerator) Intn(n int) int {
	g.calledWithN = n
	return g.randomNum
}

func TestEightBall(t *testing.T) {
	cases := []struct {
		randomNum int
		want      string
	}{
		{0, "Definitivamente no"},
		{1, "Quizás"},
		{2, "Si"},
		{3, "Absolutamente"},
		{-1, "Absolutamente"}, // default case
	}

	for _, tt := range cases {
		g := &fixedRandIntGenerator{randomNum: tt.randomNum}
		eb := EightBall{
			rand: g,
		}

		got := eb.Answer("¿Esto realmente funciona?")
		if got != tt.want {
			t.Errorf("EightBall.Answer() es %q para num %d, se queria %q",
				got, tt.randomNum, tt.want)
		}

		if g.calledWithN != 3 {
			t.Errorf("EightBall.Answer() no llamó a Intn(3) como se esperaba")
		}

	}

}
