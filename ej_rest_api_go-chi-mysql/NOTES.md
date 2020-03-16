## Pasos:

    1- crear modelos
    2- acceder a la bd (mysql)
    3- crear conectores
    4- crear webserver, handlers y routing (chi)
    5- documentacion con swagger
    6- tests
    7- makefile
    8- dockerizar proyecto

### Notas gotools:

- Ejecutar tests dirs niveles inferiores:
normal-> go test ./... | verbose-> go test -v ./... 

- Formatear codigo files directorio actual:
gofmt -w .

### Go Mods info: 

https://github.com/golang/go/wiki/Modules

Uso aqui outside gopath, alias:

- go mod init github.com/zeroidentidad/rest-chi-mysql

Arreglo detectar dependencias necesarias:

- go mod tidy

Bajar dependencias en dir de trabajo actual:

- go mod vendor

## Tareas Dir cmd:

- Init migracion: go run cmd/main/main.go