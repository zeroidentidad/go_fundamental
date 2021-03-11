// Comentario de una sola línea

/* Comentario multi-
líneas*/

/* Una etiqueta de compilación es un comentario de línea que comienza con // +build
y se puede ejecutar con el comando build -tags="foo bar" (por ser un ejemplo).
Las etiquetas de compilación se colocan antes de la cláusula del paquete, cerca
o en la parte superior del archivo, seguidas de una línea en blanco
u otros comentarios de línea. */

// +build prod, dev, test

// Una cláusula de paquete inicia cada archivo fuente.
// Main es un nombre especial que declara un ejecutable en lugar de una libreria.
package main

/* La declaración de importación declara los paquetes de la libreria
a los que se hace referencia en este archivo.*/
// ya sea por linea o agrupadas
import (
	"fmt"       // Un paquete en la libreria estándar de Go.
	"io/ioutil" // Implementa algunas funciones de utilidad de E/S.
	"log"
	m "math"   // Libreria Math con alias local m.
	"net/http" // un servidor web!
	"os"       // Funciones OS para trabajar con el sistema de archivos
	"strconv"  // Conversiones de strings.
)

// Definición de función. Main es especial. Es el punto de entrada para el
// programa ejecutable. Go usa corchetes.
func main() {
	// Println genera una línea en stdout.
	// Viene del paquete fmt.
	fmt.Println("Hola Mundo!")

	// Llama a otra función dentro de este paquete.
	beyondHello()
}

// Las funciones tienen parámetros entre paréntesis.
// Si no hay parámetros, aún se requieren paréntesis vacíos.
func beyondHello() {
	var x int // Declaración de variable. Las variables deben declararse antes de su uso.
	x = 3     // Asignación variable.
	// Las declaraciones "cortas" usan := para inferir el tipo, declarar y asignar.
	y := 4
	sum, prod := learnMultiple(x, y)        // La función devuelve dos valores a usar.
	fmt.Println("sum:", sum, "prod:", prod) // Salida simple.
	learnTypes()                            // < y minutos, aprende más!
}

/*
Las funciones pueden tener parámetros y (múltiples!) valores de retorno.
Aquí, `x`,` y` son los argumentos y `sum`,` prod` es la firma (lo que se devuelve).
Tenga en cuenta que `x` y` sum` reciben/retornan el tipo `int`.
*/
func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y // Devuelve dos valores.
}

// Algunos tipos y literales incorporados.
func learnTypes() {
	// Por lo general, una declaración corta te da lo que quieres.
	str := "Aprende Go!" // tipo string.

	s2 := `Un literal de string "sin formato" 
		puede incluir saltos de línea.` // Mismo tipo string.

	// Literal no ASCII. Go por defecto es UTF-8.
	g := 'Σ' // tipo runa, un alias int32, contiene punto de código Unicode.

	f := 3.14195 // float64, un número de punto flotante IEEE-754 de 64 bits.
	c := 3 + 4i  // complex128, representado internamente con dos float64

	// sintaxis var con inicializadores.
	var u uint = 7 // Sin signo, tamaño depende de implementación como con int.
	var pi float32 = 22. / 7

	// Sintaxis de conversión con declaración corta.
	n := byte('\n') // byte es un alias de uint8.

	// Arrays tienen tamaño fijo en tiempo de compilación.
	var a4 [4]int                    // Un array de 4 ints, inicializado todo a 0.
	a5 := [...]int{3, 1, 5, 10, 100} // Un array inicializado con un tamaño fijo de
	// cinco elementos, con valores 3, 1, 5, 10 y 100.

	// Los arrays (matrices) tienen semántica de valor.
	a4_cpy := a4                    // a4_cpy es una copia de a4, dos instancias separadas.
	a4_cpy[0] = 25                  // Solo se cambia a4_cpy, a4 permanece igual.
	fmt.Println(a4_cpy[0] == a4[0]) // false

	// Slices tienen un tamaño dinámico. Arrays y Slices tienen sus ventajas
	// pero los casos de uso de Slices son mucho más comunes.
	s3 := []int{4, 5, 9}     // Comparar con a5. No hay puntos suspensivos aquí.
	s4 := make([]int, 4)     // Asigna un slice de 4 ints, inicializado a todo 0.
	var d2 [][]float64       // Declaración solamente, nada asignado aquí.
	bs := []byte("un slice") // Sintaxis conversión de tipo.

	// Los Slices (así como los maps y los channels) tienen semántica de referencia.
	s3_cpy := s3                    // Ambas variables apuntan a la misma instancia.
	s3_cpy[0] = 0                   // Significa que ambos están actualizados.
	fmt.Println(s3_cpy[0] == s3[0]) // true

	// Debido a que son dinámicos, los Slices se pueden agregar a pedido.
	//-- aqui continuar <-
	// Para agregar elementos a un slice, se utiliza la función append().
	// El primer argumento es un slice al que estamos agregando. Comúnmente,
	// la variable array se actualiza en su lugar, como ejemplo siguiente.
	s := []int{1, 2, 3}    // El resultado es slice de longitud 3.
	s = append(s, 4, 5, 6) // Añadidos 3 elementos. Ahora tiene longitud de 6.
	fmt.Println(s)         // El slice ahora es [1 2 3 4 5 6]

	// Para agregar otro slice, en lugar de una lista de elementos atómicos,
	// podemos pasar referencia a un slice o un literal slice como este, con un
	// puntos suspensivos finales, tomar un slice y descomprimir sus elementos,
	// agregándolos al slice s.
	s = append(s, []int{7, 8, 9}...) // Segundo argumento es literal de slice.
	fmt.Println(s)                   // El slice ahora es [1 2 3 4 5 6 7 8 9]

	p, q := learnMemory() // Declara p, q de tipo puntero a int.
	fmt.Println(*p, *q)   // * sigue un puntero. Esto imprime dos ints.

	// Los mapas son un tipo de array asociativo que puede crecer dinámicamente, como
	// los tipo hash o diccionario de algunos otros lenguajes.
	m := map[string]int{"tres": 3, "cuatro": 4}
	m["uno"] = 1

	// Las variables no utilizadas son un error en Go.
	// The underscore lets you "use" a variable but discard its value.
	// El guión bajo permite "usar" una variable pero descartar su valor.
	// aka "Blank Identifier"
	_, _, _, _, _, _, _, _, _, _ = str, s2, g, f, u, pi, n, a5, s4, bs
	// Por lo general, se usa para ignorar un valor de retorno de  función
	// Por ejemplo, en un script rápido y sucio, quiza para ignorar el
	// valor de error devuelto por os.Create, y se espera que el archivo
	// siempre se creará.
	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "Así es como se escribe en un archivo")
	file.Close()

	// La salida, por supuesto, cuenta como una variable.
	fmt.Println(s, c, a4, s3, d2, m)

	learnFlowControl() // De vuelta en el flujo.
}

// Es posible, a diferencia de otros idiomas, para funciones en go
// tener valores de retorno con nombre en la firma de la funcion
// permite devolver fácilmente desde múltiples puntos,
// así como solo usar la palabra clave return, sin nada más.
func learnNamedReturns(x, y int) (z int) {
	z = x * y
	return // z está implícito aquí, porque quedo nombrado antes.
}

// Go es recolector de basura. Tiene punteros pero no aritmética de punteros.
// Puede cometerse error con un puntero nulo, pero no incrementando un puntero.
// A diferencia de C/C++, tomar y devolver dirección de variable local también es seguro.
func learnMemory() (p, q *int) {
	// Los valores devueltos con nombre p y q tienen un puntero de tipo a int.
	p = new(int) // La función incorporada new() asigna memoria.
	// El slice int asignado se inicializa a 0, p ya no es nil.
	s := make([]int, 20) // Asigna 20 ints como un solo bloque de memoria.
	s[3] = 7             // Asigna a uno de ellos.
	r := -2              // Declara otra variable local.
	return &s[3], &r     // & toma la dirección de un objeto.
}

func expensiveComputation() float64 {
	return m.Exp(10)
}

func learnFlowControl() {
	// Declaraciones IF requieren corchetes y no requieren paréntesis.
	if true {
		fmt.Println("te lo dije")
	}
	// El formato está estandarizado por el comando "go fmt".
	if false {
		// Si.
	} else {
		// Entonces.
	}

	// Switch en lugar de las sentencias if encadenadas.
	x := 42.0
	switch x {
	case 0:
	case 1, 2: // Puede tener múltiples coincidencias en un caso
	case 42:
		// Los casos no "fracasan".
		/*
			Sin embargo, hay palabra clave "fallthrough", consultar:
			  https://github.com/golang/go/wiki/Switch#fall-through
		*/
	case 43:
		// No alcanzado.
	default:
		// El caso predeterminado es opcional.
	}

	// El switch de tipo permite cambiar tipo de algo en lugar del valor
	// con apoyo de la reflexion
	var data interface{}
	data = ""
	switch c := data.(type) {
	case string:
		fmt.Println(c, "es un string")
	case int64:
		fmt.Printf("%d es un int64\n", c)
	default:
		// Todos los otros casos
	}

	// Como IF, el FOR tampoco usa parens.
	// Variables declaradas en for, e if son locales a su ámbito.
	for x := 0; x < 3; x++ { // ++ es una declaración.
		fmt.Println("iteración", x)
	}
	// x == 42 aqui.

	// For es la única declaración de ciclos en Go, tiene formas alternativas.
	for { // Bucle infinito.
		break    // hasta aqui.
		continue // No alcanzado.
	}

	// Se puede utilizar range para iterar sobre array, slice, string, map, channel.
	// range devuelve un (canal) o dos valores (array, slice, string, map).
	for key, value := range map[string]int{"one": 1, "two": 2, "tres": 3} {
		// para cada par en mapa, imprimir la clave y el valor
		fmt.Printf("key=%s, value=%d\n", key, value)
	}
	// Si solo se necesita valor, usar guión bajo como clave a omitir
	for _, name := range []string{"Tom", "Zoe", "Jesus"} {
		fmt.Printf("Hola, %s\n", name)
	}

	// Al igual que FOR,:= en sentencia IF significa declarar y asignar
	// "y" primero, luego comprueba y > x.
	if y := expensiveComputation(); y > x {
		x = y
	}
	// Los literales de función son cierres (func closures).
	xBig := func() bool {
		return x > 10000 // Referencia "x" declarada arriba de switch.
	}
	x = 99999
	fmt.Println("xBig:", xBig()) // true
	x = 1.3e3                    // Esto hace x == 1300
	fmt.Println("xBig:", xBig()) // false ahora.

	// Además, literales de función se pueden definir y llamar en línea,
	// actuando como un argumento para funcion, siempre que:
	// a) la función literal se llame inmediatamente (),
	// b) el tipo de resultado coincide con el tipo de argumento esperado.
	fmt.Println("Sumar + duplicar dos números: ",
		func(a, b int) int {
			return (a + b) * 2
		}(10, 2)) // Llamado con args 10 y 2
	// => Sumar + duplicar dos números: 24

	// Cuando se necesite, le gustará.
	goto love
love:

	learnFunctionFactory() // func retornando func es fun(3)(3)
	learnDefer()      // Un desvío rápido a una palabra clave importante.
	learnInterfaces() // Algo interesante y abundante
}

func learnFunctionFactory() {
	// Los siguientes dos son equivalentes, siendo el segundo más práctico
	fmt.Println(sentenceFactory("verano")("Un hermoso", "día!"))

	d := sentenceFactory("verano")
	fmt.Println(d("Un hermoso", "día!"))
	fmt.Println(d("Una perezosa", "tarde!"))
}

// Decoradores son comunes en otros idiomas. Lo mismo se puede hacer en Go
// con funciones literales que aceptan argumentos.
func sentenceFactory(mystring string) func(before, after string) string {
	return func(before, after string) string {
		return fmt.Sprintf("%s %s %s", before, mystring, after) // new string
	}
}

func learnDefer() (ok bool) {
	// Declaración defer empuja llamada a una función a una lista. La lista de guardados
	// las llamadas se ejecutan DESPUÉS de que retorne la función circundante.
	defer fmt.Println("declaraciones diferidas se ejecutan en orden inverso (LIFO).")
	defer fmt.Println("\nEsta línea se imprime primero porque")
	// Defer se usa comúnmente para cerrar archivo, por lo que la función que cierra el
	// archivo permanece cerca de la función que abre el archivo.
	return true
}

// Define Stringer como un tipo de interfaz con un método, String.
type Stringer interface {
	String() string
}

// Defina pair como una estructura con dos campos, ints llamados x, y.
type pair struct {
	x, y int // dependciendo de uso se agregan tags conocidas "xml, db, json, bson, gorm, etc"
}

// Define un método en tipo pair. Pair ahora implementa Stringer porque pair
// ha definido los métodos en la interfaz Stringer.
func (p pair) String() string { // p es llamado "receptor" (receiver)
	// Sprintf es otra función pública en el paquete fmt.
	// La sintaxis de puntos hace referencia a los campos de p.
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func learnInterfaces() {
	// La sintaxis de llaves es una "struct literal". Evalúa un struct
	// inicializado. La sintaxis := declara e inicializa p en esta estructura.
	p := pair{3, 4}
	fmt.Println(p.String()) // Llamar al método String de p, de tipo pair.
	var i Stringer          // Declare i de tipo interfaz Stringer.
	i = p                   // Válido porque pair implementa Stringer
	// Llamar al método String de i, de tipo Stringer. Salida igual que arriba.
	fmt.Println(i.String())

	// Funciones del paquete fmt llaman a método String para preguntar a un objeto
	// por una representación imprimible de sí mismo.
	fmt.Println(p) // Salida igual que arriba. Println llama al método String.
	fmt.Println(i) // Salida igual que el anterior..

	learnVariadicParams("estupendo", "aprendiendo", "aquí!")
}

// Las funciones pueden tener parámetros variables.
func learnVariadicParams(myStrings ...interface{}) {
	// Iterar cada valor de variadic params.
	// El guion bajo ignora el argumento de índice del array.
	for _, param := range myStrings {
		fmt.Println("param:", param)
	}

	// Pasar valor variadic como un parámetro variadic.
	fmt.Println("params:", fmt.Sprintln(myStrings...))

	learnErrorHandling()
}

func learnErrorHandling() {
	// ", ok" modismo utilizado para decir si algo funcionó/obtuvo o no.
	m := map[int]string{3: "tres", 4: "cuatro"}
	if x, ok := m[1]; !ok { // Ok será false porque 1 no está en el mapa.
		fmt.Println("no uno aqui")
	} else {
		fmt.Print(x) // x sería el valor, si estuviera en el mapa.
	}
	// Un valor de error comunica no solo "ok" sino más sobre el problema.
	if _, err := strconv.Atoi("non-int"); err != nil { // _ descarta valor
		// imprime 'strconv.ParseInt: parsing "non-int": invalid syntax'
		fmt.Println(err)
	}
	// Revisar, las interfaces un poco más tarde. Mientras tanto,
	learnConcurrency()
}

// c es un canal, un objeto de comunicación seguro para la concurrencia.
func inc(i int, c chan int) {
	c <- i + 1 // <- es operador "enviar" cuando aparece un canal a la izquierda.
}

// inc para incrementar algunos números al mismo tiempo.
func learnConcurrency() {
	// La misma función make utilizada anteriormente para crear slice. Make asigna y
	// inicializa slices, maps y channels.
	c := make(chan int)
	// Iniciar tres goroutines simultáneos. Los números se incrementarán
	// concurrentemente, tal vez en paralelo si la máquina es capaz y
	// configurado correctamente para multiprocesador. Los tres envían al mismo canal.
	go inc(0, c) // go es una declaración que inicia una nueva goroutine.
	go inc(10, c)
	go inc(-805, c)
	// Leer tres resultados del canal e imprimirlos.
	// No se sabe en qué orden llegarán los resultados.
	fmt.Println(<-c, <-c, <-c) // canal a la derecha, <- es el operador "recibir".

	cs := make(chan string)         // Otro canal, este maneja strings.
	ccs := make(chan chan string)   // Un canal de canales de strings.
	go func() { c <- 84 }()         // Inicia una goroutine solo para enviar un valor.
	go func() { cs <- "verboso" }() // Nuevamente, por cs esta vez.
	// Select tiene una sintaxis como setencia switch, pero cada caso implica
	// una operación de canal. Selecciona un caso al azar de los casos.
	// que esten listos para comunicarse.
	select {
	case i := <-c: // El valor recibido se puede asignar a una variable,
		fmt.Printf("es un %T", i)
	case <-cs: // o se puede descartar el valor recibido.
		fmt.Println("es un string")
	case <-ccs: // Canal vacío, no listo para comunicación.
		fmt.Println("no sucedió.")
	}
	// En este punto se tomó un valor de c o cs. Uno de los dos
	// goroutines iniciados arriba se ha completado, el otro permanece bloqueado.

	learnWebProgramming() // Go lo hace. Algo simple.
}

// Una sola función del paquete http inicia un servidor web.
func learnWebProgramming() {

	// El primer parámetro de ListenAndServe es la dirección TCP a escuchar.
	// El segundo parámetro es una interfaz, específicamente http.Handler.
	go func() {
		err := http.ListenAndServe(":8080", pair{}) // *: las inet disponibles
		log.Println(err)                            // no ignorar los errores
	}()

	requestServer()
}

// Hacer pair un http.Handler implementando su único método, ServeHTTP.
func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Servir datos con un método de http.ResponseWriter.
	w.Write([]byte("Aprendiste Go en X minutos!"))
}

func requestServer() {
	resp, err := http.Get("http://localhost:8080")
	log.Println(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("\nWebserver dice: `%s`", string(body))
}
