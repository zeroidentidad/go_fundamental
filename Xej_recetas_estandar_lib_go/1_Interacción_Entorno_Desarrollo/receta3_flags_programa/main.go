package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Tipo personalizado necesita para implementar
// interfaz flag.Value para poder
// usarlo en la funcion flag.Var
type ArrayValue []string

func (s *ArrayValue) String() string {
	return fmt.Sprintf("%v", *s)
}

func (a *ArrayValue) Set(s string) error {
	*a = strings.Split(s, ",")
	return nil
}

func main() {

	// Extracción de valores flag con métodos que devuelven punteros
	retry := flag.Int("retray", -1, "Define el máximo de reintentos")

	// Leer flag con la función XXXVar
	// En este caso la variable debe ser definida
	// antes del flag.
	var logPrefix string
	flag.StringVar(&logPrefix, "prefix", "", "prefijo Logger")

	var arr ArrayValue
	flag.Var(&arr, "array", "Arreglo de entrada para recorrer en iteración.")

	// Ejecutar la funcion flag.Parse para
	// que lea las flags a las variables definidas.
	// Sin esta llamada las variables permanecen vacías.
	flag.Parse()

	// Lógica de muestra no relacionada con flags
	logger := log.New(os.Stdout, logPrefix, log.Ldate)

	retryCount := 0
	for retryCount < *retry {
		logger.Println("Reintentando conexión")
		logger.Printf("Enviando array %v\n", arr)
		retryCount++
	}
}

//Ejemplo de ejecucion: go run main.go -retray 3 -prefix=ejemplo -array=1,2
