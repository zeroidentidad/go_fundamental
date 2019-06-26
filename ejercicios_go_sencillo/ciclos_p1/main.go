package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println(target)

	fmt.Println("Se ha elegido un numero entre 1 y 100.")
	fmt.Println("¿Puedes adivinarlo?")

	reader := bufio.NewReader(os.Stdin)

	adivina := 0
	for target != adivina {

		fmt.Print("Adivina el numero: ")

		entrada, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		entrada = strings.TrimSpace(entrada)
		adivina, _ = strconv.Atoi(entrada)

		if adivina < target {
			fmt.Println("Demasiado bajo")
		} else if adivina > target {
			fmt.Println("Demasiado alto")
		}

	}

	fmt.Println("Adivinaste :)")

}
