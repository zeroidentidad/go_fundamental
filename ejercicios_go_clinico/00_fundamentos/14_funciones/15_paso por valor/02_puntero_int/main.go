package main

import "fmt"

func main() {

	edad := 44
	fmt.Println(&edad) // 0xc000098000

	cambiaMe(&edad)

	fmt.Println(&edad) //0xc000098000
	fmt.Println(edad)  //24
}

func cambiaMe(z *int) {
	fmt.Println(z)  // 0xc000098000
	fmt.Println(*z) // 44
	*z = 24
	fmt.Println(z)  // 0xc000098000
	fmt.Println(*z) // 24
}
