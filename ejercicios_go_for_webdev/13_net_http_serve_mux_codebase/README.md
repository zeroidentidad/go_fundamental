# Términos sinónimo de programación web en Go

- router
- request router
- multiplexer
- mux
- servemux
- server
- http router
- http request router
- http multiplexer
- http mux
- http servemux
- http server

***

En electrónica, un [multiplexer (or mux)] (https://en.wikipedia.org/wiki/Multiplexer) es un dispositivo que selecciona una de varias señales de entrada y reenvía la entrada seleccionada en una sola línea.

El término multiplexor ha sido adoptado por la programación web para referirse al proceso de enrutamiento de solicitudes.

Un servidor web tiene solicitudes que entran en diferentes enrutadores y a través de diferentes métodos HTTP. Por ejemplo, podríamos tener estas solicitudes:

SOLICITUD # 1
  - Ruta: /gato
  - Método: GET


SOLICITUD # 2
  - Ruta: /aplicar
  - Método: GET

Solicitud n. ° 3
  - Ruta: /aplicar
  - Método: POST

En función de las solicitudes que ingresen, el servidor debe determinar cómo responder a esa solicitud; para cada solicitud que ingrese, se ejecutará un código diferente.

Se ha estado usando la palabra "servidor" pero también podría haberse estado usando la palabra "multiplexor" o "mux". El servidor, multiplexor o mux determina qué código debe ejecutarse en respuesta a cada solicitud entrante

***

ServeMux es un multiplexor de solicitud HTTP.

Un ServeMux hace coincidir la URL de cada solicitud entrante con una lista de patrones registrados y llama al controlador para el patrón que sea similar a la URL.

Los patrones nombran rutas fijas, enraizadas, como "/favicon.ico", o subárboles enraizadas, como "/images/" (tenga en cuenta la barra inclinada final).


Los patrones más largos tienen prioridad sobre los más cortos, de modo que si hay controladores registrados para "/images/" y "/images/thumbnails/", se llamará al último controlador para las rutas que comienzan por "/images/ thumbnails/" y el primero recibirá solicitudes para cualquier otra ruta en el subárbol "/images/". Tenga en cuenta que dado que un patrón que termina en una barra diagonal nombra un subárbol enraizado, el patrón "/" coincide con todas las rutas que no coinciden con otros patrones registrados, no solo la URL con Ruta == "/".

Si se ha registrado un subárbol y se recibe una solicitud nombrando la raíz del subárbol sin su barra diagonal final, ServeMux redirige esa solicitud a la raíz del subárbol (agregando la barra diagonal final). Este comportamiento se puede anular con un registro separado para la ruta sin la barra inclinada final. Por ejemplo, registrar "/images/" hace que ServeMux redirija una solicitud de "/images" a "/images/", a menos que "/images" se haya registrado por separado.

Los patrones pueden comenzar opcionalmente con un nombre de host, restringiendo las coincidencias a las URL en ese host únicamente. Los patrones específicos del host tienen prioridad sobre los patrones generales, por lo que un controlador puede registrarse para los dos patrones "/ codesearch" y "codesearch.google.com/" sin también hacerse cargo de las solicitudes de "http://www.google.com/ ".

ServeMux también se encarga de desinfectar la ruta de solicitud de URL, redirigiendo cualquier solicitud que contenga. o ... elementos o barras repetidas a una URL equivalente y más limpia.

***

# ServeMux

[http.ServeMux](https://godoc.org/net/http#ServeMux)

``` Go
type ServeMux
	func NewServeMux() *ServeMux
	func (mux *ServeMux) Handle(pattern string, handler Handler)
	func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
	func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

Cualquier valor de tipo ```*http.ServeMux``` implementa la interfaz ```http.Handler```.

Recordar, la interfaz ```http.Handler``` requiere que un tipo tenga el método ```ServeHTTP```.

```
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Lo que esto nos dice es que podemos pasar un valor de tipo ```* http.ServeMux``` a ```http.ListenAndServe```

```
func ListenAndServe(addr string, handler Handler) error
```

También puede verse en la documentación de ```*http.ServeMux``` ...

``` Go
type ServeMux
	func NewServeMux() *ServeMux
	func (mux *ServeMux) Handle(pattern string, handler Handler)
	func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
	func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

... que tenemos un método ```Handle``` adjuntado al valor de tipo ```*http.ServeMux```. Puede ver que ```Handle``` toma un patrón y un ```http.Handler```.

Podemos usar ```Handle``` de esta forma:

```
	mux := http.NewServeMux()
	mux.Handle("/", h)
	mux.Handle("/gato", g)
```

El plan general del juego:

Creamos un NewServeMux, luego adjuntamos el método ```Handle``` para establecer rutas, luego pasaremos nuestro ```* http.ServeMux``` a ```http.ListenAndServe```.


***

# DefaultServeMux

ListenAndServe inicia un servidor HTTP con una dirección y un controlador dados. El controlador suele ser nulo, lo que significa utilizar DefaultServeMux. Handle y HandleFunc agregan controladores a DefaultServeMux:

```
http.ListenAndServe(":8080", nil)
```

***

