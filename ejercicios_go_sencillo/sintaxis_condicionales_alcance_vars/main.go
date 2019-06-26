package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Escribe tu puntuacion: ")

	leer := bufio.NewReader(os.Stdin)

	//entrada, _ := leer.ReadString('\n') //se omite error de retorno con _
	entrada, err := leer.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	entrada = strings.TrimSpace(entrada)

	//puntuacion, _ := strconv.Atoi(entrada)
	puntuacion, _ := strconv.ParseFloat(entrada, 64)

	var estatus string
	if puntuacion >= 50 {
		estatus = "aprobado"
	} else {
		estatus = "suspendido"
	}

	fmt.Println("puntuacion de", puntuacion, "es", estatus)
}

/*func main() {
	puntuacion := 26
	if puntuacion == 100 {
		fmt.Println("Perfecto")
	} else if puntuacion >= 60 {
		fmt.Println("Has aprobado")
	} else {
		fmt.Println("Has suspendido")
	}
}*/
