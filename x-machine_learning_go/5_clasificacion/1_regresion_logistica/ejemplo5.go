package main

import (
	"bufio"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {

	// Crear conjunto de datos de entrenamiento y prueba

	// Abrir archivo de dataset de préstamos limpiado.
	f, err := os.Open("./clean_loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear marco de datos del archivo CSV.
	// Se inferirán los tipos de las columnas.
	loanDF := dataframe.ReadCSV(f)

	// Calcular número de elementos en cada conjunto.
	trainingNum := (4 * loanDF.Nrow()) / 5
	testNum := loanDF.Nrow() / 5
	if trainingNum+testNum < loanDF.Nrow() {
		trainingNum++
	}

	// Crear índices de subconjuntos.
	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	// Enumerar los índices de entrenamiento.
	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}

	// Enumerar los índices de prueba.
	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}

	// Crear marcos de datos del subconjunto.
	trainingDF := loanDF.Subset(trainingIdx)
	testDF := loanDF.Subset(testIdx)

	// Crear un mapa que se utilizará para escribir los datos en los archivos.
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	// Crear los archivos respectivos.
	for idx, setName := range []string{"training.csv", "test.csv"} {

		// Guardar archivo de conjunto de datos filtrado.
		f, err := os.Create("./newData/" + setName)
		if err != nil {
			log.Fatal(err)
		}

		// Crear un escritor almacenado en búfer.
		w := bufio.NewWriter(f)

		// Escribir el marco de datos como un CSV.
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}
