package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Print("Escribe tu puntuacion: ")

	leer := bufio.NewReader(os.Stdin)

	//entrada, _ := leer.ReadString('\n') //se omite error de retorno con _
	entrada, err := leer.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(entrada)
}
