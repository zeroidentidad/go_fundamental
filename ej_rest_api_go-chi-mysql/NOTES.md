## Pasos:

    1- crear modelos
    2- acceder a la bd (mysql)
    3- crear conectores
    4- crear webserver, handlers y routing (chi) 
    6- tests

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
- Init main Dir files: go run cmd/main/*.go

## Postman data test:
- URL: http://localhost:9000/smartphones
- Content-Type application/json - Raw JSON: {"name":"Pixie5","price":"2800","country_origin":"Mexico","os":"Android 6"}