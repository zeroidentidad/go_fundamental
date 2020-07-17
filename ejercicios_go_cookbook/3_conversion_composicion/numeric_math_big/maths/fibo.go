package maths

import "math/big"

// variable global para memorizar fib
var memorize map[int]*big.Int

func init() {
	// inicializar map
	memorize = make(map[int]*big.Int)
}

// Fib imprime el enésimo dígito de la secuencia de Fibonacci
// devolverá 1 para cualquier cosa < 0 también...
// se calcula de forma recursiva y usa big.Int de ser
// int64 se desbordará rápidamente
func Fib(n int) *big.Int {
	if n < 0 {
		return nil
	}

	// caso base
	if n < 2 {
		memorize[n] = big.NewInt(1)
	}

	// comprueba si se almaceno antes
	// si es así, regresar sin cálculo
	if val, ok := memorize[n]; ok {
		return val
	}

	// inicializa el mapa y luego agrega 2 valores previos de fib
	memorize[n] = big.NewInt(0)
	memorize[n].Add(memorize[n], Fib(n-1))
	memorize[n].Add(memorize[n], Fib(n-2))

	// return result
	return memorize[n]
}
