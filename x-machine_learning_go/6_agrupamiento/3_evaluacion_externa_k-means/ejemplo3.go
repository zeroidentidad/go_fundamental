package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/mash/gokmeans"
)

func main() {

	// centros agrupamiento con algoritmo k-means

	// Abrir archivo de conjunto de datos de conductores.
	f, err := os.Open("./data/fleet_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear un nuevo lector CSV.
	r := csv.NewReader(f)
	r.FieldsPerRecord = 3

	// Inicializar slice de gokmeans.Node para guardar los datos de entrada.
	var data []gokmeans.Node

	// Recorrer los registros creando slice de gokmeans.Node
	for {

		// Leer registro y verificar si hay errores de archivo.
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Omitir encabezado
		if record[0] == "Driver_ID" {
			continue
		}

		// Inicializar un punto.
		var point []float64

		// Llenar punto.
		for i := 1; i < 3; i++ {

			// Analizar valor flotante.
			val, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				log.Fatal(err)
			}

			// Agregar valor al punto.
			point = append(point, val)
		}

		// Agregar punto a los datos.
		data = append(data, gokmeans.Node{point[0], point[1]})
	}

	// Generar agrupamientos con k-means.
	success, centroids := gokmeans.Train(data, 2, 50)
	if !success {
		log.Fatal("Could not generate clusters")
	}

	// Salida de los centroides a stdout.
	fmt.Println("The centroids for our clusters are:")
	for _, centroid := range centroids {
		fmt.Println(centroid)
	}
}
