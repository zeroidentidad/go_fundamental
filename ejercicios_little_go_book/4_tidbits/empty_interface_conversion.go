package main

import "fmt"

func main() {
	suma := add(5, 5)
	fmt.Println(suma)

	wtf(3.1416)
}

func add(a interface{}, b interface{}) interface{} {
	return a.(int) + b.(int)
}

func wtf(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Printf("a is now an int and equals %d\n", a)
	case bool, string:
		fmt.Println("Enviaste un boolean o string con:", a)
	default:
		fmt.Println("keseso!")
	}
}
