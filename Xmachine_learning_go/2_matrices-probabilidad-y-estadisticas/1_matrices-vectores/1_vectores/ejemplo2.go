package main

import (
	"fmt"

	//"github.com/gonum/matrix/mat64"
	"gonum.org/v1/gonum/blas/blas64"
)

type Vector struct {
	mat blas64.Vector
	n   int
}

func main() {

	// Crear un nuevo valor vectorial.
	myvector := NewVector(2, []float64{11.0, 5.2})

	fmt.Println(myvector)
}

func NewVector(n int, data []float64) *Vector {
	if len(data) != n && data != nil {
		panic("ErrShape")
	}
	if data == nil {
		data = make([]float64, n)
	}
	return &Vector{
		mat: blas64.Vector{
			Inc:  1,
			Data: data,
		},
		n: n,
	}
}
