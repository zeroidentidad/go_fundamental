# Usar funciones en plantillas

## [documentación de la función de template](https://godoc.org/text/template#hdr-Functions)

***

## [template.FuncMap](type FuncMap map[string]interface{})

FuncMap es el tipo de mapa que define la asignación de nombres a funciones. Cada función debe tener un solo valor de retorno o dos valores de retorno de los cuales el segundo tiene un error de tipo. En ese caso, si el segundo valor de retorno (error) se evalúa como no nulo durante la ejecución, la ejecución termina y Execute devuelve ese error.

## [template.Funcs](https://godoc.org/text/template#Template.Funcs)
``` Go
func (t *Template) Funcs(funcMap FuncMap) *Template
```

***

Durante la ejecución, las funciones se encuentran en dos mapas de funciones:
- primero en la plantilla,
- luego en el mapa de funciones globales.

De manera predeterminada, no hay funciones definidas en la plantilla, pero el método Funcs se puede usar para agregarlas.

Las funciones globales predefinidas se definen en text/template.