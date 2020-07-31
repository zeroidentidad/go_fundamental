package collections

// WorkWith es la estructura en la
// que se implementará colecciones for
type WorkWith struct {
	Data    string
	Version int
}

// Filter es un filtro funcional. Toma una lista de WorkWith
// y una función WorkWith que devuelve un valor bool para cada elemento "true"
// y lo devuelve a la lista resultante
func Filter(ws []WorkWith, f func(w WorkWith) bool) []WorkWith {
	// dependiendo de los resultados, el tamaño más pequeño para el resultado
	// es len == 0
	result := make([]WorkWith, 0)
	for _, w := range ws {
		if f(w) {
			result = append(result, w)
		}
	}
	return result
}

// Map es un mapa funcional. Toma una lista de WorkWith
// y una función WorkWith que toma un WorkWith
// y devuelve un WorkWith modificado.
// El resultado final es una lista de WorkWiths modificados
func Map(ws []WorkWith, f func(w WorkWith) WorkWith) []WorkWith {
	// el resultado siempre debe ser de la misma longitud
	result := make([]WorkWith, len(ws))

	for pos, w := range ws {
		newW := f(w)
		result[pos] = newW
	}
	return result
}
