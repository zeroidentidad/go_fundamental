package main

import (
	"fmt"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	contador := 0
	for i := 0; i < b.N; i++ {
		if i%3 == 0 {
			contador += i
		} else if i%5 == 0 {
			contador += i
		}
	}
	fmt.Println(contador)
}

// ejecuta esto con comando:
// go test -bench='.*'
