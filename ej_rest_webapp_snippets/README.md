
Bajar dependencias `go mod vendor`

Recursos BD en `/sql` y `/docker-compose.yml`

Para dev TLS certificate en `/tls/generate_cert.sh`

```shell
#!/bin/bash
# dependiendo ubicación instalación del GoSDK
go run Root_GoSDK/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

Para iniciar ejecución: `./run.sh`

```shell
#!/bin/bash
export S_ADDR=":4001"
export S_DSN="webusr:passx123@tcp(localhost:3306)/pastein?parseTime=true"
export S_SECRET="z3Roh+pPbnzHbS*+9Pk8qGWhTzbpa@jf"
go run . -addr=$S_ADDR -dsn=$S_DSN -secret=$S_SECRET
```

- Ayuda configuración: `go run . -help`

Pruebas:

- TODO: `/tests`