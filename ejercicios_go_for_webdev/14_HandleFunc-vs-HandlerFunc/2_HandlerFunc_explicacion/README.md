# HandlerFunc

[http.HandlerFunc] (https://godoc.org/net/http#HandlerFunc)

```Go
tipo HandlerFunc func (ResponseWriter, *Request)
```

```Go
func (f HandlerFunc) ServeHTTP (w ResponseWriter, r * Request)
```

**Esto es bueno de saber. Probablemente no se haría esto en el código de producción.**

***

## Pregunta
¿Se podría hacer que http.Handle tome una función con esta firma: func (ResponseWriter, *Request)?