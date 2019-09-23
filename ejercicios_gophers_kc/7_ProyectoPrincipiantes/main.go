package main

import (
	"fmt"
)

func main() {
	var porcentajeExamenes int
	var porcentajeComportamiento int
	var porcentajeTrabajos int
	// Primeramente le decimos que nos diga el porcentaje
	fmt.Println("Por favor introduzca el porcentaje de Examenes")
	_, err := fmt.Scanf("%d\n", &porcentajeExamenes)
	// Comprobamos que no tenemos error mediante el err != nil significa que no tenemos error
	// _, err := fmt.Scanf("formato", & el porcentaje)
	// Para que la funcion funcine debemos poner el valor con el & para que lo reciba que lo pueda llenar correctamente
	// Cuando ponemos el _ es para decir que no queremos el valor ese y lo despreciamos.
	if err != nil {
		fmt.Println("Error al introduccir el porcentaje de Examenes")
		return
	}
	fmt.Println("Por favor introduzca el porcentaje de Comportamiento")
	_, err = fmt.Scanf("%d\n", &porcentajeComportamiento)
	if err != nil {
		fmt.Println("Error al introduccir el porcentaje de Comportamiento")
		return
	}
	fmt.Println("Por favor introduzca el porcentaje de Trabajos")
	_, err = fmt.Scanf("%d\n", &porcentajeTrabajos)
	if err != nil {
		fmt.Println("Error al introduccir el porcentaje de Trabajos")
		return
	}
	// Calculamos que la suma sea 100
	if (porcentajeComportamiento + porcentajeExamenes + porcentajeTrabajos) != 100 {
		fmt.Printf("Error ya que los porcentajes no dan 100 (%v + %v + %v != 100)\n", porcentajeExamenes, porcentajeComportamiento, porcentajeTrabajos)
		return
	}

	var numeroIteraciones int

	fmt.Println("Cuantos Usuarios quieres contabilizar")
	_, err = fmt.Scanf("%d\n", &numeroIteraciones)
	if err != nil {
		fmt.Println("Invalido el numero de usuarios a contabilizar")
		return
	}

	for i := 0; i < numeroIteraciones; i++ {
		var notaExamen int
		var notaCompotamiento int
		var notaTrabajo int

		fmt.Printf("Usuario numero: %d\n", i+1)
		fmt.Println("Introduzca la nota de Examenes")
		_, err = fmt.Scanf("%d\n", &notaExamen)
		if err != nil {
			fmt.Println("Nota del examen no valida")
			return
		}
		fmt.Println("Introduzca la nota de Comportamiento")
		_, err = fmt.Scanf("%d\n", &notaCompotamiento)
		if err != nil {
			fmt.Println("Nota de comportamiento no valida")
			return
		}
		fmt.Println("Introduzca la nota del Trabajo")
		_, err = fmt.Scanf("%d\n", &notaTrabajo)
		if err != nil {
			fmt.Println("Nota del trabajo no valida")
			return
		}
		// Hacemos que los int sean floats para que admintan coma
		notaConPorcentajeExamen := float64(porcentajeExamenes) / 100 * float64(notaExamen)
		notaConPorcentajeComportamiento := float64(porcentajeComportamiento) / 100 * float64(notaCompotamiento)
		notaConPorcentajeTrabajo := float64(porcentajeTrabajos) / 100 * float64(notaTrabajo)
		// Imprimimos con Printf usando %f para ver la coma flotante sin el formato matematico con exponenciales
		fmt.Printf("\t --> Resumen Usuario numero: %d\n", i+1)
		fmt.Printf("\t\t --> Nota Examen: %f\n", notaConPorcentajeExamen)
		fmt.Printf("\t\t --> Nota Comportamiento: %f\n", notaConPorcentajeComportamiento)
		fmt.Printf("\t\t --> Nota Trabajo: %f\n", notaConPorcentajeTrabajo)
		fmt.Printf("\t --> Resumen Nota Total Usuario: %f + %f + %f = %f\n", notaConPorcentajeExamen, notaConPorcentajeComportamiento, notaConPorcentajeTrabajo,
			notaConPorcentajeExamen+notaConPorcentajeComportamiento+notaConPorcentajeTrabajo)
	}
}
