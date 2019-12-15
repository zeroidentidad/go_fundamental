# En este paso:

Eliminaremos un usuario de mongodb.

Esto es idéntico a lo que hicimos en el último paso para GET user.

Primero obtendremos el ID de usuario de la URL

```
id := p.ByName("id")
```

A continuación, verificar que el id es un ObjectId

```
if !bson.IsObjectIdHex(id) {
	w.WriteHeader(http.StatusNotFound) // 404
	return
}
```

ObjectIdHex devuelve un ObjectId de la representación hexadecimal proporcionada.

```
	oid := bson.ObjectIdHex(id)
```

A continuación, agregar código para eliminar al usuario

```
if err := uc.session.DB("go_rest_tutorial").C("users").RemoveId(oid); err != nil {
	w.WriteHeader(404)
	return
}
```

# Ejecutar este código

1. Start server

## DELETE user mongodb

Enter en terminal

```
curl -X POST -H "Content-Type: application/json" -d '{"name":"Miss Moneypenny","gender":"female","age":27}' http://localhost:8080/user
```

```
curl http://localhost:8080/user/<enter-user-id>

```

```
curl -X DELETE http://localhost:8080/user/<enter-user-id>
```