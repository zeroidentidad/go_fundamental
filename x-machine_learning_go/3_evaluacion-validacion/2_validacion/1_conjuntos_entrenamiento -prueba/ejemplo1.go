package main

import (
	"bufio"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {

	// url data: https://archive.ics.uci.edu/ml/datasets/diabetes

	// Abrir el archivo del conjunto de datos de diabetes.
	f, err := os.Open("./data/diabetes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear un dataframe a partir del archivo CSV.
	// Se inferirán los tipos de las columnas.
	diabetesDF := dataframe.ReadCSV(f)

	// Calcule el número de elementos en cada conjunto.
	trainingNum := (4 * diabetesDF.Nrow()) / 5
	testNum := diabetesDF.Nrow() / 5
	if trainingNum+testNum < diabetesDF.Nrow() {
		trainingNum++
	}

	// Crear los índices de subconjuntos.
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

	// Crear los marcos de datos del subconjunto.
	trainingDF := diabetesDF.Subset(trainingIdx)
	testDF := diabetesDF.Subset(testIdx)

	// Cree un mapa que se utilizará para escribir los datos en los archivos.
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	// Crear los archivos respectivos.
	for idx, setName := range []string{"training.csv", "test.csv"} {

		// Guardar el archivo de conjunto de datos filtrado.
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}

		// Crea un escritor en buffer.
		w := bufio.NewWriter(f)

		// Escriba el marco de datos como CSV.
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}
