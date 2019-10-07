# Intro:
 Este ejemplo practico es para entender como se ejecutaría codigo de Go como si fuera JS en el navegador.

 En este ejemplo sencillo solo se usa el código, como sigue:

``` Go
 	variable := "Golang"
	js.Global().Get("document").Get("body").Set("innerHTML", "Hola desde la web con "+variable)
``` 


# Compilación:

``` Shell
 GOOS=js GOARCH=wasm go build -o main.wasm
``` 

# Ejecución:

## 1 - obtener copia en el proyecto actual del ejecutador js de wasm con:

``` Shell
 cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
``` 

## 2 - crear index.html como el que se encuentra en este repo e incluir el archivo del paso anterior:

``` HTML
<script src="wasm_exec.js"></script>
```

## 3 - La ejecución directa seria con agregar la herramieta goexec para probar, sin crear el programa de servidor:

``` Shell
 go get -u github.com/shurcooL/goexec
``` 

## 4 - Para servir los archivos index.html, wasm_exec.js y main.wasm en el proyecto actual desde la terminal:

``` Shell
 goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
``` 

## Para mas detalles Wiki de referencia: [go/wiki/WebAssembly](https://github.com/golang/go/wiki/WebAssembly)