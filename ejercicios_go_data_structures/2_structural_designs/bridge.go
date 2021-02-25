package main

import "fmt"

//IDrawShape interface
type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

//DrawShape struct
type DrawShape struct{}

// La estructura DrawShape tiene el método drawShape
// con coordenadas floats en X y Y
func (drawShape DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Shape")
}

//IContour interace
type IContour interface {
	drawContour(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

//DrawContour struct
type DrawContour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

//DrawContour con método drawContour dadas las coordenadas
func (contour DrawContour) drawContour(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Contour")
	contour.shape.drawShape(contour.x, contour.y)
}

//DrawContour con método resizeByFactor con factor dado
func (contour DrawContour) resizeByFactor(factor int) {
	contour.factor = factor
}

// main
func main() {

	var x = [5]float32{1, 2, 3, 4, 5}
	var y = [5]float32{1, 2, 3, 4, 5}

	var contour IContour = DrawContour{x, y, DrawShape{}, 2}

	contour.drawContour(x, y)
	contour.resizeByFactor(2)
}
