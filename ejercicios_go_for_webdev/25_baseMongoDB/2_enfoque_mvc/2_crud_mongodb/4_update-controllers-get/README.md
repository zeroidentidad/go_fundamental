# En este paso:

Obtendremos un usuario de mongodb

Primero obtendremos el id de usuario de la URL

```
id := p.ByName("id")
```

A continuación, verificaremos que el id es un ObjectId

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

Ahora crearemos un usuario vacío para almacenar los resultados en
	
```
	u := models.User{}
```

Y luego obtendremos la información del usuario

```
if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
	w.WriteHeader(404)
	return
}
```


# Ejecutar este código

1. Start server

## POST user a mongodb

Enter en terminal

```
curl -X POST -H "Content-Type: application/json" -d '{"name":"James Bond","gender":"male","age":32}' http://localhost:8080/user
```

-X abreviatura de --request
Especifica un método de solicitud personalizado para usar cuando se comunica con el servidor HTTP.

-H abreviatura de --header

-d abreviatura de --data

## GET user de mongodb

Enter en terminal

```
curl http://localhost:8080/user/<enter-user-id>

```