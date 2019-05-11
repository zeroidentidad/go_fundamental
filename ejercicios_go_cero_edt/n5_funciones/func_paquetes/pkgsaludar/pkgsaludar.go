package pkgsaludar

import "fmt"

// MeVes es para utilizar desde otro paquete
var MeVes string

var noMeVes string

//Saludar con inicial mayuscula, significa funcion exportada.
func Saludar(nombre string) {
	fmt.Println("Hola", nombre)
}

func noVisible() {
	fmt.Println("No me ven desde otro paquete")
}
