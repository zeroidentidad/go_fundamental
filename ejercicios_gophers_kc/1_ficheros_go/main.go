package main

// Si es una app --> debemos definir el package main con una función main de entrada del aplicativo,
// IMPORTANTE: Si no definimos no podrá ser ejecutado como aplicación y si hacemos un go run no ejecutará nada.
// Si es una librería --> Podemos poner el package queramos.

// Si nosotros definimos un package main debemos tener SI o SI una función main

// Definición de un import
//import "fmt"
//import "math"
// Sugar Syntax -> Tenemos dos maneras de definir un import, por convención se emplea la sintaxis de () siempre que haya más de un import
import (
	"fmt"
)

// Por cada librería que nosotros importemos emplearemos más disco en el ejecutable de salida. Por eso debemos decidir si debemos emplear cada una de las librerías sobre todo las de 3 partes.

// La función main no tiene parámetros.
func main() {
	fmt.Println()
}
