package main

import (
	"bufio"
	"fmt"
	"os"
)

type candidatos struct {
	nombre string
	votos  int
}

func main() {
	candidatoAMLO := candidatos{"AMLO", 0}
	candidatoMeade := candidatos{"Jose Meade", 0}

	votado := []string{}

	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Println(" == Sistema de votaciones ==  \n Ingrese su nombre:")
		scanner.Scan()

		usuario := scanner.Text()

		if usuario == "" {
			fmt.Println("Invalido")
			continue
		}

		tieneVoto := false

		for i := 0; i < len(votado); i++ {
			if votado[i] == usuario {
				fmt.Println("Ya votaste")
				tieneVoto = true
				break
			}
		}

		if tieneVoto {
			continue
		}

		fmt.Printf("Por quien quierese votar, %q ó %q : ", candidatoAMLO.nombre, candidatoMeade.nombre)
		scanner.Scan()

		if scanner.Text() == candidatoAMLO.nombre {
			candidatoAMLO.votos++
		} else if scanner.Text() == candidatoMeade.nombre {
			candidatoMeade.votos++
		} else {
			fmt.Println("Candidato invalido")
			continue
		}

		votado = append(votado, usuario)

		fmt.Println("Voto exitoso")

		fmt.Println("|Conteo|")
		fmt.Println(candidatoAMLO)
		fmt.Println(candidatoMeade)

	}

}
