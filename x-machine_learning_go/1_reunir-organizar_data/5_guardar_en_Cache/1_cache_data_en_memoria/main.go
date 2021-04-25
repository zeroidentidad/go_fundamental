package main

import (
	"fmt"
	"time"

	cache "github.com/patrickmn/go-cache"
)

func main() {

	// Crear caché con un tiempo de caducidad predeterminado de 5 minutos y que
	// purga los elementos caducados cada 30 segundos
	c := cache.New(5*time.Minute, 30*time.Second)

	// Colocar una clave y un valor en la caché.
	c.Set("mikey", "mivalorZero", cache.DefaultExpiration)

	// Para verificar integridad. Salida de la clave y el valor en el caché a la salida estándar.
	v, found := c.Get("mikey")
	if found {
		fmt.Printf("key: mikey, valor: %s\n", v)
	}
}
