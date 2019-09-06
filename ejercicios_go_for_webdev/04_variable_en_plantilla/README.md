# Variables template 

## [variables template](https://godoc.org/text/template#hdr-Variables)

### ASIGNAR
``` Go
{{$wisdom := .}}
```

### USO
``` Go
{{$wisdom}}
```

Un pipeline dentro de una acción puede inicializar una variable para capturar el resultado. La inicialización tiene sintaxis
 
 $variable := pipeline
 
 donde $variable es el nombre de la variable Una acción que declara una variable no produce salida.
 
Si una acción "range" inicializa una variable, la variable se establece en los elementos sucesivos de la iteración. Además, un "range" puede declarar dos variables, separadas por una coma:
 
  range $index, $element := pipeline

 en cuyo caso $index y $element se establecen en los valores sucesivos del índice de array/slice o clave y elemento del map, respectivamente. Tener en cuenta que si solo hay una variable, se le asigna el elemento; Esto es opuesto a la convención en las cláusulas range de Go.
 
El alcance de una variable se extiende a la acción "end" de la estructura de control ("if", "with" o "range") en la que se declara, o al final de la plantilla si no existe dicha estructura de control. Una invocación de plantilla no hereda variables desde el punto de su invocación.
 
 When execution begins, $ is set to the data argument passed to Execute, that is, to the starting value of dot.

 Cuando comienza la ejecución, $ se establece en el argumento de datos pasado a Ejecutar, es decir, en el valor inicial de punto.