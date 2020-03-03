package main

import (
	"fmt"
	"log"

	"github.com/lytics/anomalyzer"
)

func main() {

	// Inicializar valor AnomalyzerConf con configuraciones
	// de qué métodos de detección de anomalías se usara.
	conf := &anomalyzer.AnomalyzerConf{
		Sensitivity: 0.1,
		UpperBound:  5,
		LowerBound:  anomalyzer.NA, // ignorar el límite inferior
		ActiveSize:  1,
		NSeasons:    4,
		Methods:     []string{"diff", "fence", "highrank", "lowrank", "magnitude"},
	}

	// Crear serie de tiempo de observaciones periódicas como slice de flotantes.
	// Esto podría provenir de una base de datos o archivo, como en ejemplos anteriores.

	ts := []float64{0.1, 0.2, 0.5, 0.12, 0.38, 0.9, 0.74} // ejemplo hardcode

	// Crear nuevo anomalyzer basado en los valores y
	// la configuración de series de tiempo existentes.
	anom, err := anomalyzer.NewAnomalyzer(conf, ts)
	if err != nil {
		log.Fatal(err)
	}

	// Anomalyzer analizará el valor en referencia a valores preexistentes en la serie
	// y generará una probabilidad de que el valor sea anómalo.
	prob := anom.Push(15.2)
	fmt.Printf("Probability of 15.2 being anomalous: %0.2f\n", prob)

	prob = anom.Push(0.43)
	fmt.Printf("Probability of 0.33 being anomalous: %0.2f\n", prob)
}
