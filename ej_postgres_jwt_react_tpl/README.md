# Go-React template
![](https://awebytes.files.wordpress.com/2020/08/logo.png)

Este es un ejemplo boilerplate/starter para un proyecto en Golang (con postgres) y React.

*Se utiliza gin framework en go.

## Uso

Clonar o descargar este repositorio

Utilizar scripts en el archivo [server/db/.psql](./server/db/.psql) para configurar la base de datos.

Ingresar las credenciales de conexion a DB en el archivo [server/config/config.go](./server/config/config.go) 

Navegar al directorio **server** ya teniendo las configuraciones previas.

```bash
> cd server
> go run .
```

Esto iniciará el servidor Go.

Para iniciar la aplicación React, navegar hasta el directorio **client**

```bash
> cd client
> yarn install
> yarn start
```
### EndPoints

* /session [GET]

* /register [POST]
```js
       { 
         name String,
         email String,
         password String
       }
```

* /login [POST]
```js
       { 
         email String,
         password String
       }
```

* /createReset [POST]
```js
       {
         email String
       }
```

* /resetPassword [POST]
```js
       {
         id Int,
         password String,
         confirm_password String
       }
```

## Rutas

* /login

* /register

* /session

* /createReset

* /resetPassword


## Contribuir
Pull requests son bienvenidas. Para cambios importantes, primero abrir un issue para discutir qué le gustaría cambiar.