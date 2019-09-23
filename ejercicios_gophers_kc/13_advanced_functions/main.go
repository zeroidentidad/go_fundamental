package main

import (
	"fmt"
	"math"
)

type FuncionesCalculadora func(float64, float64) float64

func computar(funcion FuncionesCalculadora, a float64, b float64) float64 {
	return funcion(a, b)
}

func dameFuncionPorSigno(signo string) FuncionesCalculadora {
	switch signo {
	case "+":
		return func(a, b float64) float64 {
			return a + b
		}
	case "-":
		return func(a, b float64) float64 {
			return a - b
		}
	case "*":
		return func(a, b float64) float64 {
			return a * b
		}
	case "^":
		return math.Pow
	default:
		return func(a, b float64) float64 {
			return a
		}
	}
}

func main() {
	fmt.Println(dameFuncionPorSigno("+")(1, 2))
	fmt.Println(computar(dameFuncionPorSigno("+"), 1, 2))
	fmt.Println(computar(dameFuncionPorSigno("-"), 1, 2))
	fmt.Println(computar(dameFuncionPorSigno("^"), 2, 3))
}
