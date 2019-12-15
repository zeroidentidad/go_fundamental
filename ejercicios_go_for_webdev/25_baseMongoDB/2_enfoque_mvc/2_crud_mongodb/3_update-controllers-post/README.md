# No ejecute este código todavía

Creamos un ObjectId usando el paquete bson.

Hacemos esto en controllers/user.go en func CreateUser

`` `
// crear bson ID
u.Id = bson.NewObjectId ()

`` `

En segundo lugar, almacenamos al usuario en mongodb.

Hacemos esto en controllers/user.go en func CreateUser

`` `
uc.session.DB ("go-web-dev-db"). C ("users"). Insertar (u)