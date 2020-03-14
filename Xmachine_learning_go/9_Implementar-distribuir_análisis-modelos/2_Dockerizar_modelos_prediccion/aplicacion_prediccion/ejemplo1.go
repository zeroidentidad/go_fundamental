package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// ModelInfo incluye la información sobre el modelo que sale del entrenamiento.
type ModelInfo struct {
	Intercept    float64           `json:"intercept"`
	Coefficients []CoefficientInfo `json:"coefficients"`
}

// CoefficientInfo incluye información sobre un coeficiente de modelo particular.
type CoefficientInfo struct {
	Name        string  `json:"name"`
	Coefficient float64 `json:"coefficient"`
}

// PredictionData incluye los datos necesarios para hacer una predicción
// y codifica la predicción de salida.
type PredictionData struct {
	Prediction      float64          `json:"predicted_diabetes_progression"`
	IndependentVars []IndependentVar `json:"independent_variables"`
}

// IndependentVar incluye información y un valor para una variable independiente.
type IndependentVar struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func main() {

	// Declarar las banderas de directorio de entrada y salida.
	inModelDirPtr := flag.String("inModelDir", "", "The directory containing the model.")
	inVarDirPtr := flag.String("inVarDir", "", "The directory containing the input attributes.")
	outDirPtr := flag.String("outDir", "", "The output directory")

	// Analizar banderas de línea de comandos.
	flag.Parse()

	// Cargar el archivo del modelo.
	f, err := ioutil.ReadFile(filepath.Join(*inModelDirPtr, "model.json"))
	if err != nil {
		log.Fatal(err)
	}

	// Descomprimir información del modelo.
	var modelInfo ModelInfo
	if err := json.Unmarshal(f, &modelInfo); err != nil {
		log.Fatal(err)
	}

	// Avanzar sobre los archivos en la entrada.
	if err := filepath.Walk(*inVarDirPtr, func(path string, info os.FileInfo, err error) error {

		// Omitir cualquier directorio.
		if info.IsDir() {
			return nil
		}

		// Abrir cualquier archivo.
		f, err := ioutil.ReadFile(filepath.Join(*inVarDirPtr, info.Name()))
		if err != nil {
			return err
		}

		// Descomponer las variables independientes.
		var predictionData PredictionData
		if err := json.Unmarshal(f, &predictionData); err != nil {
			return err
		}

		// Hacer la predicción.
		if err := Predict(&modelInfo, &predictionData); err != nil {
			return err
		}

		// Marshal los datos de predicción.
		outputData, err := json.MarshalIndent(predictionData, "", "    ")
		if err != nil {
			log.Fatal(err)
		}

		// Guardar la salida ordenada en un archivo.
		if err := ioutil.WriteFile(filepath.Join(*outDirPtr, info.Name()), outputData, 0644); err != nil {
			log.Fatal(err)
		}

		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

// Predict hace una predicción basada en la entrada JSON.
func Predict(modelInfo *ModelInfo, predictionData *PredictionData) error {

	// Inicializar valor de predicción a la intersección.
	prediction := modelInfo.Intercept

	// Crear mapa de coeficientes de variables independientes.
	coeffs := make(map[string]float64)
	varNames := make([]string, len(modelInfo.Coefficients))
	for idx, coeff := range modelInfo.Coefficients {
		coeffs[coeff.Name] = coeff.Coefficient
		varNames[idx] = coeff.Name
	}

	// Crear un mapa de los valores de las variables independientes.
	varVals := make(map[string]float64)
	for _, indVar := range predictionData.IndependentVars {
		varVals[indVar.Name] = indVar.Value
	}

	// Recorrer las variables independientes.
	for _, varName := range varNames {

		// Obtener coeficiente.
		coeff, ok := coeffs[varName]
		if !ok {
			return fmt.Errorf("Could not find model coefficient %s", varName)
		}

		// Obtener el valor de la variable.
		val, ok := varVals[varName]
		if !ok {
			return fmt.Errorf("Expected a value for variable %s", varName)
		}

		// Añadir a la predicción.
		prediction = prediction + coeff*val
	}

	// Agregar predicción a los datos de predicción.
	predictionData.Prediction = prediction

	return nil
}
