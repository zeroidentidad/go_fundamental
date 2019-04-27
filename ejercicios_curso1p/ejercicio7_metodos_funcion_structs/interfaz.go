// Declarar interface llamada speaker con un método llamado speak. 
// Declarar struct llamado ingles que represente a una persona que hable Inglés.
// Declarar struct llamado chino que represente a una persona que hable chino.
// Implementar interface speaker para cada struct usando un valor y strings: "Hello World" y "你好世界"
// Declarar variable del tipo speaker y asignarle una dirección de tipo Inglés y llamar el método.
// Y para Chino.


package main

import "fmt"

// speaker implementa lo que dice una persona
type speaker interface {
	speak()
}

// inglés representa a una persona que habla inglés

type ingles struct{}

// speak implementa la interface speaker
func (ingles) speak() {
	fmt.Println("Hello World")
}

// chino representa a una persona que habla chino

type chino struct{}

// speak implementa la interface speaker
func(chino) speak() {
	fmt.Println("你好世界")
}

func main() {
	//Declarar una variable para interface
	var sp speaker

	// Asignarle un valor a la interface y vamos a llamar al método (de la interface)
	var i ingles
	sp = i
	sp.speak()

	// Chino
	var c chino
	sp = c
	sp.speak()

	// Crear nuevos valores y llamar la función
	decirHola(new(ingles))
	decirHola(&chino{})
}

// decirHola abstrae la función de hablar

func decirHola(sp speaker) {
	sp.speak()
}


	







