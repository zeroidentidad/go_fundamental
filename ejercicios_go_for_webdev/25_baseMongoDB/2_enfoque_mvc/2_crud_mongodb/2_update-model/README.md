No ejecutes este código

Simplemente haciendo modificaciones: un proceso de varios pasos.

IMPORTANTE:
Actualizar declaraciones de importación para importar paquetes desde la ubicación correcta!

En este paso:

MongoDB representa documentos JSON en formato codificado en binario llamado BSON detrás de escena. BSON extiende el modelo JSON para proporcionar tipos de datos adicionales y ser eficiente para codificar y decodificar en diferentes lenguajes.

Actualizaremos nuestro modelo de usuario para cambiar el tipo del campo Id para que sea un bson.

Agregar esto a models/user.go

```
type User struct {
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
	Id     bson.ObjectId `json:"id" bson:"_id"`
}

```