package main

import (
	"bufio"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {

	// Crear conjunto de datos de entrenamiento y pruebas (parte 3)

	// Abrir archivo de dataset
	f, err := os.Open("./data/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear un marco de datos a partir del archivo CSV.
	// Se inferirán los tipos de las columnas.
	advertDF := dataframe.ReadCSV(f)

	// Calcular el número de elementos en cada conjunto.
	trainingNum := (4 * advertDF.Nrow()) / 5
	testNum := advertDF.Nrow() / 5
	if trainingNum+testNum < advertDF.Nrow() {
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

	// Create the subset dataframes.
	// Crear los marcos de datos del subconjunto.
	trainingDF := advertDF.Subset(trainingIdx)
	testDF := advertDF.Subset(testIdx)

	// Cree un mapa que se utilizará para escribir los datos en los archivos.
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	// Crear los archivos respectivos.
	for idx, setName := range []string{"training.csv", "test.csv"} {

		// Guardar el archivo de conjunto de datos filtrado.
		f, err := os.Create("./newData/" + setName)
		if err != nil {
			log.Fatal(err)
		}

		// Crea un escritor protegido.
		w := bufio.NewWriter(f)

		// Escribir el marco de datos como un CSV.
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}
