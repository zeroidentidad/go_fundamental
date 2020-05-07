package bridge

// Abstraccion Interface
type DispositivoInterface interface {
	ConectarInternet() string
	SetConexion(Conexion)
}

// Abstraccion Abstracta
type Dispositivo struct {
	conexion Conexion
}

func (d *Dispositivo) SetConexion(conexion Conexion) {
	d.conexion = conexion
}

// Abstraccion Refinada
type Telefono struct {
	numero string
	*Dispositivo
}

func (t *Telefono) ConectarInternet() string {
	return "Teléfono N° " + t.numero + " conectado a internet mediante " + t.conexion.Conectar()
}

// Abstraccion Refinada
type Tablet struct {
	*Dispositivo
}

func (t *Tablet) ConectarInternet() string {
	return "Tablet conectada a internet mediante " + t.conexion.Conectar()
}

// Implementador Interface
type Conexion interface {
	Conectar() string
}

// Implementador Concreto
type Red4G struct{}

func (r *Red4G) Conectar() string {
	return "4G"
}

// Implementador Concreto
type RedWiFi struct{}

func (r *RedWiFi) Conectar() string {
	return "WiFi"
}
