package config

import "fmt"

func GetDir() string {
	fmt.Println("Ubicación a compatir [enter para omitir]: ")
	var ruta string
	fmt.Scanf("%s", &ruta)

	if ruta == "" {
		fmt.Println("------------------------------------------")
		fmt.Println("Usando valor 'DirShared' de config.json")
		return DirShared()
	} else {
		fmt.Println("------------------------------------------")
		fmt.Println("Ok")
		return ruta
	}
}
