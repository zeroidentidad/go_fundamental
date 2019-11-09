# Paso 1

Creado ```func getUser``` se coloca en un nuevo archivo, session.go. Este refactor permite usar el mismo código en el index y bar.

```
func getUser(w http.ResponseWriter, req *http.Request) user 
```

  
 ## IMPORTANTE
Ahora que se tiene dos archivos go del paquete main, no se puede usar "go run main.go" para ejecutar la aplicación. Eso solo pide un archivo go: main.go. Se debe usar "go build" y luego ejecutar el ejecutable creado o ejecutar con "go run * .go"
 
# Paso 2

Creado ```func signup``` y eliminado el código de registro de ```func index```. Se agregó un nuevo campo para la contraseña a la estructura del usuario (user struct).

```
func signup(w http.ResponseWriter, req *http.Request)
```

# Paso 3

Created ```func alreadyLoggedIn``` and put it on the session.go page. This refactor allows us to use the same code in bar and signup.

Creado ```func alreadyLoggedIn``` se coloca en la página session.go. Este refactor permite usar el mismo código en bar y signup.

```
func alreadyLoggedIn(req *http.Request) bool
```