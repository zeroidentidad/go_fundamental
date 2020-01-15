package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("goroutines demo\n")

	// run func en background
	go calculate()

	index := 0

	// run en background
	go func() {
		for index < 6 {
			fmt.Printf("go func() index= %d \n", index)
			var result float64 = 2.5 * float64(index)
			fmt.Printf("go func() result= %.2f \n", result)

			time.Sleep(500 * time.Millisecond)
			index++
		}
	}()

	// run en background
	go fmt.Printf("go printed in the background\n")

	// press ENTER to exit
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

}

func calculate() {
	i := 12
	for i < 15 {
		fmt.Printf("calculate()= %d \n", i)
		var result float64 = 8.2 * float64(i)
		fmt.Printf("calculate() result= %.2f \n", result)

		time.Sleep(500 * time.Millisecond)
		i++
	}
}
