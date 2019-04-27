 // declarar estructura para guardar info de usuario
// declarar funcion que crea el valor y regresa apuntadores de este tipo y un error como valor
// llamar esta funcion desde main y regresar el valor

// hacer una segunda llamada a la funcion, pero ignorando el valor y solo probando el error como valor

package main

import "fmt"

type usuario struct { 
	nombre string
	email string 
}

// func nuevoUsuario crea y regresa apuntadores de valores tipo usuario

func nuevoUsuario() (*usuario, error){
	return &usuario{ "Jesus", "jesusfake@mail.com" }, nil
}

func main(){
	//Crear valor tipo usuario

	u, err:=nuevoUsuario()
	if err!=nil{
		fmt.Println(err)
		return
	}

	//Imprimir valor
	fmt.Println(*u)

	// 2do llamado a la funcion y que solo cheque el error en el regreso

	_, err = nuevoUsuario()
	if err != nil {
		fmt.Println(err)
		return
	}
}