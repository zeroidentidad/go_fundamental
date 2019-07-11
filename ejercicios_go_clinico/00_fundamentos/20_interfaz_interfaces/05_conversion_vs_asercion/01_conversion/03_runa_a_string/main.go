package main

import "fmt"

func main() {
	var x rune = 'a' // rune es un alias para int32; normalmente omitido en esta sentencia
	var y int32 = 'b'
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(string(x))
	fmt.Println(string(y))
	// conversion: rune to string
}
