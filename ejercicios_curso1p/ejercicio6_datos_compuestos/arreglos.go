
//declarar array de 5 cadenas con 0, y otros 5 con los valores del anterior

//asignar el primero al segundo y mostrarlo

package main

import "fmt"

func main(){

	// arreglo vacio
	var nombres[5] string

	// arreglo con valores

	chicas:= [5] string {"Fulana 1", "Fulana 2", "Fulana 3", "Fulana 4", "Fulana 5"}

	// asignar el arreglo de los chicas al arreglo nombres

	nombres = chicas

	// mostrar cada valor del arreglo y su posicion en el array

	for i, nombre := range nombres {
		fmt.Println("Nombre: ", nombre, " Posicion: ", i)
	} 

	/* Version curso: */
	//Declarar arreglos de strings que guarden los nombres
	var nombresx[5]string
	// Declarar un arreglo pre-llenado con nombres de amigos.
	amigos:= [5]string{"Raquel", "Isabel", "Fernando", "Enrique", "José"}
	// Asignar el arreglo de los amigos al arreglo nombres
	nombresx = amigos
	// Mostrar cada uno de los valores string y la dirección de cada uno en el 2ndo arreglo
	fmt.Println("Version curso:")
	for h, nombrea := range nombresx {
		fmt.Println(nombrea, &nombresx[h])
	}

}