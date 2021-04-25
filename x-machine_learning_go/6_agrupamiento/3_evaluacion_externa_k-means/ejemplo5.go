package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/gonum/floats"
)

func main() {

	// evaluar agrupamientos generados parte 2 obteniendo distancia media centroides

	// Abrir archivo conjunto de datos conductores.
	f, err := os.Open("./data/fleet_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear marco de datos a partir del archivo CSV.
	driverDF := dataframe.ReadCSV(f)

	// Extraer la columna de distancia.
	distances := driverDF.Col("Distance_Feature").Float()

	// clusterOne y clusterTwo mantendrán los valores para el trazado.
	var clusterOne [][]float64
	var clusterTwo [][]float64

	// Rellenar los grupos con datos.
	for i, speed := range driverDF.Col("Speeding_Feature").Float() {
		distanceOne := floats.Distance([]float64{distances[i], speed}, []float64{50.05, 8.83}, 2)
		distanceTwo := floats.Distance([]float64{distances[i], speed}, []float64{180.02, 18.29}, 2)
		if distanceOne < distanceTwo {
			clusterOne = append(clusterOne, []float64{distances[i], speed})
			continue
		}
		clusterTwo = append(clusterTwo, []float64{distances[i], speed})
	}

	// Producir métricas dentro del clúster.
	fmt.Printf("\nCluster 1 Metric: %0.2f\n", withinClusterMean(clusterOne, []float64{50.05, 8.83}))
	fmt.Printf("\nCluster 2 Metric: %0.2f\n", withinClusterMean(clusterTwo, []float64{180.02, 18.29}))
}

// withinClusterMean calcula la distancia media entre puntos en un grupo y el centroide del grupo.
func withinClusterMean(cluster [][]float64, centroid []float64) float64 {

	// meanDistance mantendrá resultado.
	var meanDistance float64

	// Pasar sobre los puntos en el grupo.
	for _, point := range cluster {
		meanDistance += floats.Distance(point, centroid, 2) / float64(len(cluster))
	}

	return meanDistance
}
