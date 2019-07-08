package main

import "fmt"

var a = "esto guardado en la variable a"           // alcanse package
var b, c string = "guardado en b", "guardado en c" // alcanse package
var d string                                       // alcanse package

func main() {

	d = "guardado en d" // declaracion arriba; asignacion aqui; alcanse package
	var e = 42          // alcance funcion - variables posteriores tienen el mismo alcanse package:
	f := 43
	g := "guardado en g"
	h, i := "guardado en h", "guardado en i"
	j, k, l, m := 44.7, true, false, 'm' // single quotes
	n := "n"                             // double quotes
	o := `o`                             // back ticks

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c - ", c)
	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)
	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)
	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)
	fmt.Println("m - ", m)
	fmt.Println("n - ", n)
	fmt.Println("o - ", o)
}
