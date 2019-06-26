package main

import (
	"fmt"
)

func main() {
	/*	var ancho, largo, area float64
		ancho = 7.5
		largo = 8.3
		area = ancho * largo
		fmt.Println(area/10.0, "cajas cesped necesarias")

		ancho = 6.5
		largo = 7.7
		area = ancho * largo
		fmt.Println(area/10.0, "cajas cesped necesarias")*/

	/*fmt.Println("Calcula un Tercio: ", 1.0/3.0)

	fmt.Printf("Calcula un tercio: %0.2f\n", 1.0/3.0)*/

	//fmt.Printf("El costo de los %s es de %d euros cada uno.\n ", "balones", 4)

	/*fmt.Printf("Un float: %f\n", 4.3432)
	fmt.Printf("Un integer: %d\n", 32)
	fmt.Printf("Un string: %s\n", "Noticia")
	fmt.Printf("Un booleano: %t\n", true)
	fmt.Printf("Valores: %v %v %v\n", 4.3, "\t", false)
	fmt.Printf("Valores: %#v %#v %#v\n", 4.3, "\t", false)
	fmt.Printf("Tipos: %T %T %T\n", 4.3, "\t", false)
	fmt.Printf("Signo tanto por ciento: %%\n")*/

	//fmt.Printf("%f cajas necesarias\n", 3.454940393023)

	/*fmt.Printf("%12s | %s\n", "Producto", "Costo en Euros")

	fmt.Println("-------------------------------")

	fmt.Printf("%12s | %2d\n", "Abono", 73)
	fmt.Printf("%12s | %2d\n", "Herramientas", 43)
	fmt.Printf("%12s | %2d\n", "madera", 4)*/

	/*fmt.Printf("%8.2f\n", 1.494044949420594)
	fmt.Printf("%.2f\n", 3.872902654132456)*/

	var ancho, largo, area float64
	ancho = 7.5
	largo = 8.3
	area = ancho * largo
	fmt.Printf("%.2f cajas de cesped necesarias\n", area/10.0)

	ancho = 6.5
	largo = 7.7
	area = ancho * largo
	fmt.Printf("%.2f cajas de cesped necesarias\n", area/10.0)
}
