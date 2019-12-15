# Instalar Mongo

# Go get driver mongo

```
go get gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson
```

# En este paso:

No ejecutes este código

Simplemente haciendo actualizaciones: un proceso de varios pasos.

Necesitaremos una sesión de mongo para usar en los métodos CRUD.

Necesitamos que nuestro UserController tenga acceso a una sesión de mongo.

Agreguemos esto a controllers/user.go

```
UserController struct {  
    session *mgo.Session
}
```

Y ahora agregar esto a controllers/user.go

```
func NewUserController(s *mgo.Session) *UserController {  
    return &UserController{s}
}
```

Y ahora agregar esto a main.go

```
func getSession() *mgo.Session {
	// Connect local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
```

y esto

```
uc := controllers.NewUserController(getSession())  
```

1. Ingresar en la terminal

```
curl http://localhost:8080/user/1
```