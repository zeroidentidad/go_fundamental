package main

import (
	"fmt"
	"reflect"
)

func main() {

	var miInt int
	fmt.Println(reflect.TypeOf(&miInt))

	var miFloat float64
	fmt.Println(reflect.TypeOf(&miFloat))
}
