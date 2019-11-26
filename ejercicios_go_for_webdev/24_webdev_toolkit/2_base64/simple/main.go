package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "El amor no es más que una canción para cantar. El miedo es la forma en que morimos. Puedes hacer sonar las montañas. O hacer llorar a los ángeles. Aunque el pájaro está en el ala. Y es posible que no sepas por qué. Vamos gente ahora. Sonríe a tu hermano. Todos se juntan. Intenta amar uno al otro ahora."

	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println("\nLong original:", len(s))
	fmt.Println("\nLong codificada:", len(s64))
	fmt.Println("\nOriginal:", s)
	fmt.Println("\nCodificada:", s64)
}
