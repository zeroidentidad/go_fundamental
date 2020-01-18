package main

import "log"

func main() {
	/* archivo := "prueba.txt"
	f, err := os.Open(archivo)

	defer f.Close()

	if err != nil {
		fmt.Println("error abriendo el archivo")
		os.Exit(1)
	}   */

	ejemploPanic()
}

func ejemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error que generó Panic \n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("se encontro el valor de 1")
	}
}
