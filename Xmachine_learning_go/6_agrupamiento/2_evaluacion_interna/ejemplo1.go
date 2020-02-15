package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

type centroid []float64

func main() {

	// Evaluacion interna

	// Abrir para recuperar del archivo CSV.
	irisFile, err := os.Open("./data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Crear un marco de datos a partir del archivo CSV.
	irisDF := dataframe.ReadCSV(irisFile)

	// Definir nombres de las tres especies separadas contenidas en el archivo CSV.
	speciesNames := []string{
		"Iris-setosa",
		"Iris-versicolor",
		"Iris-virginica",
	}

	// Crear un mapa para contener información del centroide.
	centroids := make(map[string]centroid)

	// Filtrar el conjunto de datos en tres marcos de datos separados,
	// cada uno correspondiente a una de las especies de Iris.
	for _, species := range speciesNames {

		// Filtrar conjunto de datos original.
		filter := dataframe.F{
			Colname:    "species",
			Comparator: "==",
			Comparando: species,
		}
		filtered := irisDF.Filter(filter)

		// Calcular la media de las características.
		summaryDF := filtered.Describe()

		// Colocar la media de cada dimensión en el centroide correspondiente.
		var c centroid
		for _, feature := range summaryDF.Names() {

			// Omitir columnas irrelevantes
			if feature == "column" || feature == "species" {
				continue
			}
			c = append(c, summaryDF.Col(feature).Float()[0])
		}

		// Agregar centroid al map
		centroids[species] = c
	}

	// Verificacion de integridad, salida de centroides
	for _, species := range speciesNames {
		fmt.Printf("%s centroid: %v\n", species, centroids[species])
	}
}
