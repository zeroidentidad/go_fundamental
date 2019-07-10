package main

import "fmt"

type animal struct {
	sonido string
}

type perro struct {
	animal
	amigable bool
}

type gato struct {
	animal
	molesto bool
}

func main() {
	fido := perro{animal{"woof"}, true}
	fifi := gato{animal{"meow"}, true}
	shadow := perro{animal{"woof"}, true}
	mascotas := []interface{}{fido, fifi, shadow}
	fmt.Println(mascotas)
}
