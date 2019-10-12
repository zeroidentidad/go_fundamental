# Handler

* Esta es una cosa importante a saber *

[http.Handler] (https://godoc.org/net/http#Handler)
``` Go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

***

# Server

[http.ListenAndServe] (https://godoc.org/net/http#ListenAndServe)
``` Go
func ListenAndServe(addr string, handler Handler) error
```

[http.ListenAndServeTLS] (https://godoc.org/net/http#ListenAndServeTLS)
``` Go
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
```

* Tener en cuenta que las dos funciones anteriores toman un controlador *

***

# Request

Consultar [http.Request] (https://godoc.org/net/http#Request) en la documentación.
 
Con *la mayoría de los comentarios y algunos de los campos* eliminados:

```go 
type Request struct {
    Method string
    URL *url.URL
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Foo": {"Bar", "two"},
	//	}
    Header Header
    Body io.ReadCloser
    ContentLength int64
    Host string
    // Este campo solo está disponible después de llamar a ParseForm.
    Form url.Values
    // Este campo solo está disponible después de llamar a ParseForm.
    PostForm url.Values
    MultipartForm *multipart.Form
    // RemoteAddr permite registrar servidores HTTP y otro software
    // la dirección de red que envió la solicitud, generalmente para loggin.
    RemoteAddr string
}
```

Consultar también el índice que muestra el tipo [Request]() del paquete http.

- Algunas cosas interesantes que se puede hacer con un request:

## Recuperar datos de URL y formulario

`` http.Request``` es una estructura. Tiene los campos ```Form``` y ```PostForm```. Si leemos la documentación sobre estos, vemos:

```
    // El formulario contiene los datos analizados del formulario, incluida la URL
    // parámetros de consulta del campo y los datos del formulario POST o PUT.
    // Este campo solo está disponible después de llamar **ParseForm**.
    // El cliente HTTP ignora Form y usa Body en su lugar. Form url.Values

    // PostForm contiene los datos de formulario analizados de POST, PATCH, o PUT en parámetros del cuerpo.
    // Este campo solo está disponible después de llamar **ParseForm**.
    // El cliente HTTP ignora PostForm y utiliza Body en su lugar. PostForm url.Values

```

Si vemos **ParseForm**,

```go goc (r * Request) ParseForm () error ```

vemos que este es un método adjunto a una * http.Request.

***

Si nos fijamos en **FormValue ***

``` go func (r *Request) FormValue(key string) string```

vemos que este es un método adjunto a una *http.Request. FormValue devuelve el primer valor para el componente con nombre de la consulta. Los parámetros de cuerpo POST y PUT tienen prioridad sobre los valores de cadena de consulta de URL. FormValue llama a ParseMultipartForm y ParseForm si es necesario e ignora cualquier error devuelto por estas funciones. Si la clave no está presente, FormValue devuelve la cadena vacía. Para acceder a múltiples valores de la misma clave, llama a ParseForm y luego inspecciona Request.Form directamente.


***

## Ver el método HTTP

El tipo ```http.Request``` es una estructura que tiene un campo ```Method```.

***

## Ver valores de URL

El tipo ```http.Request``` es una estructura que tiene un campo ```URL```. Tener en cuenta que el tipo es una ```*url.URL```

Vistazo al tipo ```url.URL```

``` go
type URL struct {
    Scheme     string
    Opaque     string    // encoded opaque data
    User       *Userinfo // username and password information
    Host       string    // host or host:port
    Path       string
    RawPath    string // encoded path hint (Go 1.5 and later only; see EscapedPath method)
    ForceQuery bool   // append a query ('?') even if RawQuery is empty
    RawQuery   string // encoded query values, without '?'
    Fragment   string // fragment for references, without '#'
}
```

***

## Trabajar con el encabezado HTTP

El tipo ```http.Request``` es una estructura que tiene un campo ```Header```.

De la documentación sobre la estructura ```http.Header```, vemos que:

```
type Header map[string][]string
```

***

# Respuesta

[http.ResponseWriter] (https://godoc.org/net/http#ResponseWriter)
``` Go
type ResponseWriter interface {
    // Header devuelve el mapa de encabezado que será enviado por
    // WriteHeader. Cambiar el encabezado después de una llamada a
    // WriteHeader (o Write) no tiene efecto
    Header() Header

    // Write escribe los datos en la conexión como parte de una respuesta HTTP.
    //
    // Si aún no se ha llamado a WriteHeader, Write llama
    // WriteHeader (http.StatusOK) antes de escribir los datos. Si el encabezado
    // no contiene una línea de Content-Type, Write agrega un conjunto de Content-Type
    // al resultado de pasar los 512 bytes iniciales de datos escritos a
    // DetectContentType.
    Write([]byte) (int, error)

    // WriteHeader envía un encabezado de respuesta HTTP con código de estado.
    // Si WriteHeader no se llama explícitamente, la primera llamada a Write
    // activará un WriteHeader implícito (http.StatusOK).
    // Por lo tanto, las llamadas explícitas a WriteHeader se utilizan principalmente para
    // envíar códigos de error.
    WriteHeader (int)
}
```

***

## Establecer un encabezado de respuesta

Un ```http.ResponseWriter```` tiene un método ```Header()``` que devuelve un ```http.Header```.

Mirar la documentación para ```http.Header```

``` Go
type Header map[string][]string

```

Mirar los métodos que se adjuntan para el tipo ```http.Header```

``` go
type Header
func (h Header) Add(key, value string)
func (h Header) Del(key string)
func (h Header) Get(key string) string
func (h Header) Set(key, value string)
func (h Header) Write(w io.Writer) error
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error
```

Podemos establecer encabezados para una respuesta como esta:

``` Go
res.Header().Set("Content-Type", "text/html; charset=utf-8")
```