package estructuras

type Usuario struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

type Respuesta struct {
	Status  string  `json:"status"`
	Data    Usuario `json:"data"`
	Mensaje string  `json:"mensaje"`
}
