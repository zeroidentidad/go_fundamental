package main

import (
	"bytes"
	"testing"
)

//go test -run TestCountWords -v
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("palabra1 palabra2 palabra3 palabra4\n")

	exp := 4
	res := counter(b, false)

	if res != exp {
		t.Errorf("Expected %d, got %d.\n", exp, res)
	}
}

//go test -run TestCountLines -v
func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("palabra1 palabra2 palabra3\nlinea2\nlinea3 palabra1")

	exp := 3
	res := counter(b, true)

	if res != exp {
		t.Errorf("Expected %d, got %d.\n", exp, res)
	}
}
