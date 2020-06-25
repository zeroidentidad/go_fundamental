package models

/*Tweet captura de Body, el mensaje que llega en request*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
