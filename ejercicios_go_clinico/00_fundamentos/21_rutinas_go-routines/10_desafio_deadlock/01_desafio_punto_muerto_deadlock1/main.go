package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

// Esto resulta en un punto muerto.
// ¿Puedes determinar por qué?
// ¿Y qué harías para arreglarlo?
