package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// cadena simple
	var str string
	str = "Hello world"
	fmt.Println(str)

	// Concatenando
	str1 := "hello"
	str2 := "world"
	str3 := str1 + str2
	fmt.Println(str3)

	str4 := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(str4)

	// Cadena a numérico
	fmt.Println("----demo String To Numeric----")
	str_val1 := "5"
	str_val2 := "2.8769"

	var err error
	var int_val1 int64
	int_val1, err = strconv.ParseInt(str_val1, 10, 32) // base10, 64 size
	if err == nil {
		fmt.Println(int_val1) //
	} else {
		fmt.Println(err)
	}

	var float_val2 float64
	float_val2, err = strconv.ParseFloat(str_val2, 64) // 64 size
	if err == nil {
		fmt.Println(float_val2) //
	} else {
		fmt.Println(err)
	}

	// numérico a cadena
	fmt.Println("----demo numeric to string----")
	num1 := 8
	num2 := 5.872

	var str_num1 string
	str_num1 = fmt.Sprintf("%d", num1)
	fmt.Println(str_num1)

	var str_num2 string
	str_num2 = fmt.Sprintf("%f", num2)
	fmt.Println(str_num2)

	// Analizar cadenas
	fmt.Println("----demo String Parser----")
	data := "Berlin;Amsterdam;London;Tokyo"
	fmt.Println(data)
	cities := strings.Split(data, ";")
	for _, city := range cities {
		fmt.Println(city)
	}

	// Longitud de datos de cadena
	fmt.Println("----demo String Length----")
	str_data := "welcome to go"
	len_data := len(str_data)
	fmt.Printf("len=%d \n", len_data)

	// copiar datos
	fmt.Println("----demo copy data----")
	sample := "hello world, go!"
	fmt.Println(sample)
	fmt.Println(sample[0:len(sample)])               // copiar todo
	fmt.Println(sample[:5])                          // copiar 5 caracteres
	fmt.Println(sample[3:8])                         // copiar caracteres del índice 3 al 8
	fmt.Println(sample[len(sample)-5 : len(sample)]) // 5 caracteres de copia desde final de la cadena

	// Upper y Lower Case caracteres
	fmt.Println("----demo upper/lower characters----")
	message := "Hello World, GO!"
	fmt.Println(message)
	fmt.Println(strings.ToUpper(message)) // upper chars
	fmt.Println(strings.ToLower(message)) // lower chars
}
