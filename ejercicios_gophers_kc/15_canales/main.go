package main

import (
	"fmt"
)

func suma(a, b int, canal chan int) {
	canal <- a + b
}

func main() {

	canal := make(chan int, 50)
	go suma(10, 10, canal)
	go suma(20, 20, canal)
	resultado1, resultado2 := <-canal, <-canal
	// TH []-------------20---------- Go1
	// TH [20,40]----------------------- Go2
	fmt.Println(resultado1, resultado2)

	canalBuffer := make(chan int, 2)
	canalBuffer <- 1
	// TH -------------1---------- TH
	canalBuffer <- 2
	// TH -------------1---------- <- 2 TH Error si no empleamos un canal con buffer
	close(canalBuffer)

	for resultado := range canalBuffer {
		fmt.Println(resultado)
	}

}
