package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/gonum/floats"
)

type centroid []float64

func main() {

	// Evaluacion interna con Silhouette Coefficient

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

	// Crear mapa para contener el marco de datos filtrado para cada clúster.
	clusters := make(map[string]dataframe.DataFrame)

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

		// Agregar marco de datos filtrado al mapa de clústeres.
		clusters[species] = filtered

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

	// Convertir las etiquetas en slice de cadenas y crear slice
	// de nombres de columnas flotantes por conveniencia.
	labels := irisDF.Col("species").Records()
	floatColumns := []string{
		"sepal_length",
		"sepal_width",
		"petal_length",
		"petal_width",
	}

	// Pasar sobre los registros acumulando el promedio de coeficiente silhouette.
	var silhouette float64

	for idx, label := range labels {

		// a almacenará el valor acumulado para a.
		var a float64

		// Recorrer los puntos de datos en el mismo clúster.
		for i := 0; i < clusters[label].Nrow(); i++ {

			// Obtener el punto de datos para la comparación.
			current := dfFloatRow(irisDF, floatColumns, idx)
			other := dfFloatRow(clusters[label], floatColumns, i)

			// Agregar a a.
			a += floats.Distance(current, other, 2) / float64(clusters[label].Nrow())
		}

		// Determinar el otro grupo más cercano.
		var otherCluster string
		var distanceToCluster float64
		for _, species := range speciesNames {

			// Omitir el clúster que contiene el punto de datos.
			if species == label {
				continue
			}

			// Calcular la distancia al grupo desde el grupo actual.
			distanceForThisCluster := floats.Distance(centroids[label], centroids[species], 2)

			// Reemplazar el clúster actual si es relevante.
			if distanceToCluster == 0.0 || distanceForThisCluster < distanceToCluster {
				otherCluster = species
				distanceToCluster = distanceForThisCluster
			}
		}

		// b almacenará el valor acumulado para b.
		var b float64

		// Recorrer los puntos de datos en el otro grupo más cercano.
		for i := 0; i < clusters[otherCluster].Nrow(); i++ {

			// Obtener el punto de datos para la comparación.
			current := dfFloatRow(irisDF, floatColumns, idx)
			other := dfFloatRow(clusters[otherCluster], floatColumns, i)

			// Añadir a b.
			b += floats.Distance(current, other, 2) / float64(clusters[otherCluster].Nrow())
		}

		// Agregar al promedio de coeficiente silhouette.
		if a > b {
			silhouette += ((b - a) / a) / float64(len(labels))
		}
		silhouette += ((b - a) / b) / float64(len(labels))
	}

	// Salida del promedio de coeficiente silhouette final a stdout.
	fmt.Printf("\nAverage Silhouette Coefficient: %0.2f\n\n", silhouette)
}

// dfFloatRow recupera un slice de valores flotantes de un DataFrame
// en el índice dado y para los nombres de columna dados.
func dfFloatRow(df dataframe.DataFrame, names []string, idx int) []float64 {
	var row []float64
	for _, name := range names {
		row = append(row, df.Col(name).Float()[idx])
	}
	return row
}
