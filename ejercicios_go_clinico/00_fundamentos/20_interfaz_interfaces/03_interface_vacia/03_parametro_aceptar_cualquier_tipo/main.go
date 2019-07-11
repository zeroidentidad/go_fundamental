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

func specs(a interface{}) {
	fmt.Println(a)
}

func main() {
	fido := perro{animal{"woof"}, true}
	fifi := gato{animal{"meow"}, true}
	specs(fido)
	specs(fifi)
}
