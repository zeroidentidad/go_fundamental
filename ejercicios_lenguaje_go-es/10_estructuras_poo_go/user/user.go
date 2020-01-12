package user

import "time"

type Users interface {
	Login()
	esAdmin() bool
	AltaUsuario(id int, nombre string, fechaalta time.Time, status bool, categoria string)
}

type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
	Categoria string
}

type Administrador struct {
	Usuario
}

type Supervisor struct {
	Usuario
}

func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool, categoria string) {
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.Status = status
	this.Categoria = categoria
}

func (this *Usuario) Login() {

	/* intrucciones de login */
}

func (this *Usuario) esAdmin() bool {
	if this.Categoria == "A" {
		return true
	} else {
		return false
	}
}

/*func DevuelveCategoria(u Users) string{
	return format("Soy un %s",u.Categoria)
}*/
