package main

import "fmt"

func main() {
	var valor interface{} = 7
	fmt.Println(valor.(int) + 6)
}
