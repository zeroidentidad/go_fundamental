package main

import "fmt"

func main() {
	var a int
	var b int8

	a = 2389
	b = 44

	//casting
	c := a + int(b)

	fmt.Println(c)
	fmt.Printf("c es de tipo %T", c)

	// prioridad aritmetica: (), %, *, /, +, -
	d := 6 + 6*(6-6)

	fmt.Println(d)

	nom := "jesus"
	edad := 26

	fmt.Printf("Hola %s tu edad es %d \n", nom, edad)

	// Valores cero, por defecto
	var nombre string
	var numero int
	var valido bool
	fmt.Println(nombre, numero, valido)
}

/*
Referencia: https://golang.org/ref/spec#Types
*/

// Booleano : true , false

// Numerico
/*
Eje:

uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
*/

// String : cadena:= "texto"

// Array
/*
Eje:

[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))
*/

// Slice
/*
Eje:

SliceType = "[" "]" ElementType

make([]T, length, capacity)
make([]int, 50, 100)
new([100]int)[0:50]
*/

// Struct

/*
Eje:

// An empty struct.
struct {}

// A struct with 6 fields.
struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}

// A struct with four embedded fields of types T1, *T2, P.T3 and *P.T4
struct {
	T1        // field name is T1
	*T2       // field name is T2
	P.T3      // field name is T3
	*P.T4     // field name is T4
	x, y int  // field names are x and y
}
*/

// Pointer

/*
Eje:

PointerType = "*" BaseType .
BaseType    = Type .
*Point
*[4]int
*/

// Function : https://golang.org/ref/spec#Function_types

// Interface : https://golang.org/ref/spec#Interface_types

// Map : https://golang.org/ref/spec#Map_types

// Channel : https://golang.org/ref/spec#Channel_types
