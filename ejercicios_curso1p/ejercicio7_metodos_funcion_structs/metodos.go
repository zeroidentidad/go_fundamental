
//declarar metodos para los structs

// Declarar un struct que representa jugador de futbol (nombre, goles, partidos)
// Declarar metodo que calcule el promedio de goles de jugador (f: goles / partidos)

//Declarar slice para inicializar slice de varios jugadores
//Iterar slice mostrando los jugadores y su promedio de goles

package main

import "fmt"

type jugador struct {
	nombre string
	goles int
	partidos int
}

// promedio de goles

func (j *jugador) promedio() float64{
	return float64(j.goles)/float64(j.partidos)
}

func main(){
	jugadores:=[]jugador{
		{"carlos", 20, 60 },
		{"adrian", 10, 50 },
		{"pedro", 5, 20 },
	}

	//mostrar promedio
	for _, jugador := range jugadores{
		fmt.Println((&jugador).promedio())
	} 
}