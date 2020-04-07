# Go modules

Los módulos permiten controlar el uso de versiones de los paquetes en Go.

Esta documentación sirve para la versión ≥ 1.13.

Cuando creamos archivos `.go` fuera del $GOPATH y este archivo no necesita de paquetes externos, Go crea una ruta false de $GOPATH para ejecutarlo.

Usa SemVer (versión semántica) para el control de versiones. [https://semver.org/lang/es/](https://semver.org/lang/es/)

Revisa la integridad de los paquetes por medio de los checksum.

## go.mod

El archivo go.mod tiene la siguiente estructura:

    module repo-server/user/package
    
    go 1.XX
    
    require (
    	repo-server/user/package Vx.y.z
    	repo-server/user/package Vx.y.z // indirect
    )
    
    replace repo-server/user/package => ../path/in/your/pc/package
    exclude repo-server/user/package Vx.y.z

Ejemplo:

    module github.com/zeroidentidad/sql-utils
    
    go 1.13
    
    require (
    	github.com/satori/uuid V1.0.1
    	github.com/labstack/echo V3.3.1 // indirect
    )
    
    replace github.com/satori/uuid => ../pruebas/uuid
    exclude github.com/labstack/echo V3.3.0

## SemVer

El versionamiento se usa así:

> Vx.y.z-prerelease.x.y.z

Siempre inicia con la letra `v` en minúscula, luego la versión `major`, `minor` y `patch`.

`x major` identifica una versión estable, que tiene compatibilidad sólo con las versiones minor y patch de dicha versión. Es decir, la versión 1.11.3 es compatible con la versión 1.7.0, pero no con la versión 2.0.0

`y minor` identifica una versión que tiene nuevas características dentro de la versión estable.

`z patch` son fix o mejoras de rendimiento en una versión estable.

`prerelease.x.y.z` go mod lo usa como la fecha de una versión que no tiene el tag de la versión. 

Ejemplo de prerelease normal: v1.0.3-rc0, v1.0.3-rc1, etc

Fuente: [https://semver.org/lang/es/](https://semver.org/lang/es/)

## Crear un módulo

    $ mkdir ~/mis-repos/pruebas
    $ cd ~/mis-repos/pruebas
    $ go mod init github.com/user/imagic
    $ vim imagic.go // Crear el código necesario
    $ go get -u [repo-server/user/package]

Lo anterior crea la carpeta `pruebas` en `/home/tu-usuario`, se ubica en dicha carpeta, luego inicia el control de módulos con `go mod init ...` y luego se crea el archivo que pertenece al paquete. Si se requeren paquetes, se importan utilizando `go get`

Al iniciar el control de módulos, se crea el archivo `go.mod` y en dicho archivo se irá creando los `require` necesarios para el control de versiones. También se ha creado el archivo `go.sum` que contiene el `checksum` de cada paquete con su respectiva versión.

Los `require` que tienen la anotación `// indirect` son aquellos que se han importado usando `go get` pero que aún no se usan en el paquete que estamos creando o que uno de los paquetes de terceros que estamos usando utiliza dicho paquete en su go.mod y estamos sobreescribiendo la versión que usa el paquete de terceros.

Si no se está usando el paquete marcado con `// indirect` cuando se compile con `go build` el `// indirect` desaparecerá.

## Listar módulos

    $ go list

Muestra el módulo que está trabajando. En este documento de ejemplo, mostrará `github.com/zeroidentidad/imagic`

    $ go list all // Todos los módulos que están en $GOPATH
    $ go list -m all // Todos los módulos del paquete actual
    $ go list -m -versions [repo-server/user/package] // Muestra todas las versiones disponibles del paquete solicitado
    $ go mod verify // Verifica que la versión de un módulo no ha sido modificada

## Caché de los paquetes

En go modules todos los paquetes se descargan una sola vez y se mantienen en una sola carpeta: `$GOPATH/pkg/mod`

Si se modifica algo de esta carpeta el  `go mod verify` mostrará un error de verificación.

Ojo: La verificación no se hace al compilar.

Si se ha modificado algo, `se debe borrar todo el paquete` y volver a hacer el `go get -u` de lo contrario siempre mostrará un error de verificación.

## Ordenar los paquetes del require

Si se deja de usar un paquete, el archivo `go.mod` no lo quita del `require` tampoco lo deja como `// indirect` así que para poder limpiar este archivo, se debe utilizar la instrucción:

    $ go mod tidy

La anterior instrucción revisa qué paquetes ya no se están usando y los quita. ***Se recomienda usarlo siempre!***

## Marcar la versión de un paquete

Para marcar la versión de un paquete usamos los tag de git.

    $ git tag -a v1.11.3 -m "Mejora rendimiento en imagenes"

El anterior código marca el actual commit.

Si necesitamos marcar un commit específico, se indica el hash de dicho commit:

    $ git tag -a v1.11.3 9a2bb33 -m "Mejora rendimiento en imagenes"

Para listar los tag:

    $ git tag // muestra los tag del repo
    $ git tag -l "v1.8.*" // lista todos los tag que inicien con v1.8.
    $ git show v1.4 // muestra información detallada del commit de dicho tag

## Uso de versiones específicas

Para usar una versión específica dentro de los import de go, se usa la siguiente sintaxis:

    package mypackage
    
    import "github.com/zeroidentidad/imagic/v2"
    
    ...

La versión que está usando el anterior código es la versión 2 de imagic, pero no especifica cual versión minor o patch, para eso está el archivo go.mod

El uso de `go get -u` descarga la ultima versión estable de un paquete, si se quiere utilizar una versión específica de un paqute se puede hacer de la siguiente manera:

    $ go get -u github.com/zeroidentidad/imagic@v1.3.4

## Uso de versiones mayores o diferentes

Se puede usar diferentes versiones del paquete importando en el archivo .go el paquete y la versión:

    package main
    
    import (
    				"rsc.io/quote"
    				quoteV2 "rsc.io/quote/v2"
    )
    
    func main() {
    				quote.Proverb()
    				quoteV2.Proverb()
    }

## Commit sin versionamiento

Cuando un paquete no tiene versionamiento con los tag, go modules usa el timestamp y el hash del commit para identificar la versión.

    module github.com/zeroidentidad/imagic
    
    go 1.13
    
    require (
    				golang.org/x/tools v0.0.0-20190114222345-bf090417da8b
    )

## Consultar módulos

Se pueden usar difrentes fomas de llamar la versión de un módulo:

- Por versión específica: @v1.2.3
- Por prefijo: @v1
- Última versión: @latest
- Por commit específico: @c9af223a
- Por rama: @master
- Por comparación: @>=v1.7.2

Reglas de comparación:

- El más cercano.
- Pre-releases tienen menor precedencia.
- El mar cercano. (muy importante)

Gracias a las comparaciones se puede obtener un paquete indicando la versión que se necesita:

    $ go get github.com/zeroidentidad/imagic@v1.6.1

En el archivo `go.mod` queda especificada la versión a usar.

Si se usa @master, trae la última versión sin semver, es decir el último commit con fecha y hash.

`Latest` trae la última versión con semver.

Si vas a usar comparaciones `> < >= <=`, debes usar comillas ya que esos comparadores son caracteres especiales en bash.

    $ go get "github.com/zeroidentidad/imagic@<v1.7.0"

Si se coloca en `go.mod` una comparación, al compilar el archivo se actualizará con una versión específica para garantizar el uso del paquete.

## Argumentos de go mod

Existen varios comandos para administrar nuestros paquetes y nuestro archivo go.mod:

- Buscar el uso del modulo:

        $ go mod why [modulo]
        $ go mod why -m [modulo] // Busca quienes están usando el paquete

- Gráfico de uso

        $ go mod graph

- Ayuda de un argumento

        $ go mod help [arg]

- Cambiar el nombre de un módulo

        $ go mod edit -module [modulo]

- Cambiar la versión de go

        $ go mod edit -go 1.xx

- Agregar una dependencia al archivo go.mod

        $ go mod edit -require [modulo]

- Borrar una dependencia del archivo go.mod

        $ go mod edit -droprequire [modulo]

- Agrega una exclusión. Sirve para excluir una versión específica del paquete.

        $ go mod edit -exclude [modulo]

- Borrar una exclusión

        $ go mod edit -dropexclude [modulo]

- Reemplazar una dependencia con una local

        $ go mod edit -replace [modulo]=[modulo/local]

- Borrar un reemplazo

        $ go mod edit -dropreplace [modulo]

- Listar

        $ go mod edit -print
        $ go mod edit -json

- Descargar una dependencia (pero no la usa)

        $ go mod download [modulo]

## Vendor

Podemos tener una copia de los módulos de los que depende nuestro módulo en una carpeta interna llamada `vendor`

Para esto usamos el comando

    $ go mod vendor

Ese comando crea la carpeta vendor y hace una copia de todos los módulos de los que depende nuestro paquete.

También se crea una carpeta modules.txt la cual permite controlar las versiones de dicha carpeta.

Si se requiere compilar con la carpeta vendor, se usa:

    $ go build -mod=vendor .
    $ go run -mod=vendor .

## Orden correcto para hacer un commit

    $ go mod tidy
    $ go test ./...
    $ git add go.mod go.sum a.go b.go c.go
    $ git commit -m "Cambios correctos"
    $ git tag v1.0.0
    $ git push origin v1.0.0

## Notas adicionales

Los archivos go.mod y go.sum deben ser tenidos en cuenta en el control de versiones de git. Esto, según el blog de go: [https://blog.golang.org/using-go-modules](https://blog.golang.org/using-go-modules)