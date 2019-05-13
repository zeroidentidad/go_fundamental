package main

import "./pkganimales"

func main() {
	var p pkganimales.Perro
	var g pkganimales.Gato

	p.Nombre = "Lulu"
	g.Nombre = "Loco"

	// AdoptarPerro(p)
	// AdoptarGato(g)

	AdoptarMascota(p)
	AdoptarMascota(g)
}

/* func AdoptarPerro(p pkganimales.Perro) {
	p.Comunicarse()
}

func AdoptarGato(g pkganimales.Gato) {
	g.Comunicarse()
}
*/

func AdoptarMascota(m pkganimales.Mascota) {
	m.Comunicarse()
}
