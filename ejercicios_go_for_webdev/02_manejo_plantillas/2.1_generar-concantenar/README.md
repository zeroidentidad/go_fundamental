# Técnicas para aprender

- concatenar
- CLI pipeline: salida a un archivo con >

# Código que usaremos de la libreria estándar

## [os.Create](https://godoc.org/os#Create)
Esto nos permite crear un archivo.
``` Go
func Create (name string) (*File, error)
```

***

## [defer](https://golang.org/ref/spec#Defer_statements)
La palabra clave defer nos permite diferir la ejecución de una declaración hasta que regrese la función en la que hemos colocado la declaración defer.

***

## [io.Copy](https://godoc.org/io#Copy)
Esto nos permite copiar desde una fuente a un destino.
``` Go
func Copy (dst Writer, src Reader) (written int64, err error)
```

## [strings.NewReader](https://godoc.org/strings#NewReader)
NewReader devuelve una nueva lectura de Reader desde s.
``` Go
func NewReader (s string) *Reader
```

## [os.Args](https://godoc.org/os#pkg-variables)
Args es una variable del paquete os. Los argumentos contienen los argumentos de la línea de comandos, comenzando con el nombre del programa.
``` Go
var Args []string
```