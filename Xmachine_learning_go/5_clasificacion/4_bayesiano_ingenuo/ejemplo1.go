package main

import (
	"fmt"
	"log"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/filters"
	"github.com/sjwhitworth/golearn/naive"
)

func main() {

	// Leer dataset de entrenamiento de préstamos en "instancias" de golearn.
	trainingData, err := base.ParseCSVToInstances("./data/training.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// Inicializar un nuevo clasificador Naive Bayes.
	nb := naive.NewBernoulliNBClassifier()

	// Ajustar el clasificador Naive Bayes.
	nb.Fit(convertToBinary(trainingData))

	// Leer en dataset de prueba de préstamos en "instancias" de golearn.
	testData, err := base.ParseCSVToInstances("./data/test.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// Hacer predicciones
	predictions, _ := nb.Predict(convertToBinary(testData))

	// Generar una matriz de confusión.
	cm, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		log.Fatal(err)
	}

	// Obtener la precisión.
	accuracy := evaluation.GetAccuracy(cm)
	fmt.Printf("\nAccuracy: %0.2f\n\n", accuracy)
}

// convertToBinary utiliza la funcionalidad integrada de golearn
// para convertir las etiquetas a un formato de etiqueta binario.
func convertToBinary(src base.FixedDataGrid) base.FixedDataGrid {
	b := filters.NewBinaryConvertFilter()
	attrs := base.NonClassAttributes(src)
	for _, a := range attrs {
		b.AddAttribute(a)
	}
	b.Train()
	ret := base.NewLazilyFilteredInstances(src, b)
	return ret
}
