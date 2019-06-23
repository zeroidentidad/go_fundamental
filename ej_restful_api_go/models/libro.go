package models

//Libro structura de datos base
type Libro struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
	Anio   int    `json:"anio"`
}
