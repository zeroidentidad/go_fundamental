package main

import "fmt"

func main() {
	terminado := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func(u string) {
			fmt.Println(u)
			terminado <- true
		}(v)
	}

	// espera a que todas las goroutines se completen antes de salir
	for range values {
		<-terminado
	}
}

/*
Para vincular el valor actual de v a cada cierre a medida que se inicia,
uno debe modificar el bucle interno para crear una nueva variable en cada iteración.
una forma es pasar la variable como argumento al cierre.

En este ejemplo, el valor de v se pasa como un argumento a la función anónima.
Ese valor es accesible dentro de la función como la variable u.

SOURCE:
https://golang.org/doc/faq#closures_and_goroutines
*/
