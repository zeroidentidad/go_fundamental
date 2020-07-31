package collections

import "strings"

// LowerCaseData realiza un ToLower a la
// cadena de datos de un WorkWith
func LowerCaseData(w WorkWith) WorkWith {
	w.Data = strings.ToLower(w.Data)
	return w
}

// IncrementVersion incrementa version de WorkWiths
// Version
func IncrementVersion(w WorkWith) WorkWith {
	w.Version++
	return w
}

// OldVersion devuelve un cierre
// que valida la versión es mayor que
// la cantidad especificada
func OldVersion(v int) func(w WorkWith) bool {
	return func(w WorkWith) bool {
		return w.Version >= v
	}
}
