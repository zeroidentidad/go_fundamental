package main

import "fmt"

func main() {
	nums := []int8{52, 63, 47, 5, 5, 3, 7, 6, 100, 2, 47, -5}
	maximo, minimo := MaxyMin(nums)
	fmt.Println("Máximo:", maximo)
	fmt.Println("Mímino:", minimo)
}

//MaxyMin: eje return multiple
func MaxyMin(numeros []int8) (max int8, min int8) {
	//var max, min int8
	for _, v := range numeros {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return //max, min
}
