package main

import (
	"fmt"
	"time"
)

func main() {
	tTime := time.Date(2020, time.January, 9, 10, 5, 2, 0, time.Local)

	// El formateo se realiza con el uso del valor de referencia.
	// Jan 2 15:04:05 2006 MST
	fmt.Printf("tTime es: %s\n", tTime.Format("2006/1/2"))

	fmt.Printf("La hora es: %s\n", tTime.Format("15:04"))

	//Los formatos predefinidos podrían usarse
	fmt.Printf("La hora es: %s\n", tTime.Format(time.RFC1123))

	// El formato admite el relleno de espacio
	//solo para días en la versión Go 1.9.2 >
	fmt.Printf("tTime es: %s\n", tTime.Format("2006/1/_2"))

	// El relleno cero se hace agregando 0
	fmt.Printf("tTime es: %s\n", tTime.Format("2006/01/02"))

	//La fracción con ceros a la izquierda usa 0s (n 0)
	fmt.Printf("tTime es: %s\n", tTime.Format("15:04:05.00"))

	//La fracción sin ceros a la izquierda usa 9s (n 9)
	fmt.Printf("tTime es: %s\n", tTime.Format("15:04:05.999"))

	// Agregar formato agrega el tiempo formateado al búfer dado
	fmt.Println(string(tTime.AppendFormat([]byte("Se acabó el tiempo: "), "03:04PM")))
}
