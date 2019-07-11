package main

import "fmt"

func main() {
	terminado := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		v := v // crear una nueva 'v'.
		go func() {
			fmt.Println(v) // aqui ya no se indica warning de go vet
			terminado <- true
		}()
	}

	// Espera a que todas las goroutines se completen antes de salir
	for range values {
		<-terminado
	}
}

/*
Aún más fácil es simplemente crear una nueva variable,
usar un estilo de declaración que puede parecer extraño pero funciona bien en Go.

SOURCE:
https://golang.org/doc/faq#closures_and_goroutines
*/
