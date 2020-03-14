package main

import (
	"fmt"
)

func main() {
	//tipo{elementos} //COMPOSITE LITERAL
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Printf("%T", x)
}

//Usamos slices para agrupar VALORES del mismo TIPO
