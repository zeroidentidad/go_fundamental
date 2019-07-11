package main

import "fmt"

func main() {
	terminado := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v)
			terminado <- true
		}()
	}

	// espera a que todas las goroutines se completen antes de salir
	for range /*_ =*/ values {
		<-terminado
	}
}

/*
Puede surgir cierta confusión al utilizar cierres con concurrencia.

Uno podría esperar erróneamente ver a, b, c como la salida.
Lo que probablemente verás en cambio es c, c, c. Esto es porque
cada iteración del bucle usa la misma instancia de la variable v,
Así que cada cierre comparte esa única variable. Cuando el cierre se ejecuta,
imprime el valor de v en el momento en que se ejecuta fmt.Println,
pero v puede haber sido modificado desde que se lanzó el goroutine.
Para ayudar a detectar este y otros problemas antes de que ocurran,
Ejecutar: go vet.

SOURCE:
https://golang.org/doc/faq#closures_and_goroutines
*/
