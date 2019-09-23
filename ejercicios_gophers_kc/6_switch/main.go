package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	if runtime.GOOS == "darwin" {
		fmt.Println("macOS")
	}
	if runtime.GOOS == "linux" {
		fmt.Println("linux")
	}
	if runtime.GOOS == "windows" {
		fmt.Println("windows")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
		fmt.Println(os)
	case "linux":
		fmt.Println("linux")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println("No soportado para exportacion")
		fmt.Println(os)
	}
	ahora := time.Now().Weekday()
	switch time.Saturday {
	case ahora + 0:
		fmt.Println("Hoy")
	case ahora + 1:
		fmt.Println("Mañana")
	case ahora + 2:
		fmt.Println("Pasado Mañana")
	default:
		fmt.Println("Relajate aun queda tiempo")
	}

}
