// range itera sobre la mayoria de estructuras de datos

package main

import "fmt"

func main(){
	numeros := [] int{2,4,6} //slice
	suma := 0

	for _, numero := range numeros {
		suma += numero
	}

	fmt.Println("suma: ", suma)

	for i, numero := range numeros {
		if numero == 4 {
			fmt.Println("Indice: ", i)
		}
	}

	//algo := map[string]string{"a":auto,"b":bebe}


}