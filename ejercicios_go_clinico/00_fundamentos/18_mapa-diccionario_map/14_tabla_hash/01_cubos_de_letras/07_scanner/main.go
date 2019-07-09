package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Una fuente de entrada artificial.
	const input = "El hombre en la arena: No es el crítico quien cuenta; ni el hombre que señala cómo tropieza el hombre fuerte, o dónde el hacedor de las obras podría haberlos hecho mejor. El crédito pertenece al hombre que está realmente en la arena, cuya cara está manchada por el polvo, el sudor y la sangre; quien se esfuerza valientemente; quién se equivoca, quién se queda corto una y otra vez, porque no hay esfuerzo sin error y fallas; pero quien realmente se esfuerza por hacer las obras; Quien conoce grandes entusiasmos, las grandes devociones; que se gasta en una causa digna; quién, en el mejor de los casos, sabe al final el triunfo del gran logro, y quién, en el peor, si falla, al menos falla al atreverse, de modo que su lugar nunca será con esas almas frías, reprimidas y tímidas que no conocen la victoria ni la derrota. - Theodore Roosevelt"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Set de la funcion split para la operación de escaneo.
	scanner.Split(bufio.ScanWords)
	// Contar la palabras.
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
