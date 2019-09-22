package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	//pdf := gofpdf.New("P", "mm", "A4", "") <- config por default
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	// <- ultimo parametro sin Dir de fuentes personalizadas

	//indentación de acuerdo al page size:
	w, h := pdf.GetPageSize()
	//xIndent := w/10.0
	fmt.Printf("ancho=%v, alto=%v \n", w, h)

	pdf.AddPage()

	/* ::Cosas de texto basicas:: */
	pdf.MoveTo(0, 0) // opcional
	pdf.SetFont("times", "B", 24)
	//pdf.Cell(40, 10, "Hello, world")
	_, lineHeight := pdf.GetFontSize()
	pdf.SetTextColor(255, 0, 0) // rgb
	pdf.Text(0, lineHeight, "Contenido de prueba")

	pdf.MoveTo(0, lineHeight*2.0)
	pdf.SetFont("arial", "", 18)
	pdf.SetTextColor(100, 100, 100)
	_, lineHeight = pdf.GetFontSize()

	pdf.MultiCell(0, lineHeight*1.5, "Sed eu mauris nulla. In vitae laoreet nisi, eget maximus orci. Sed vel aliquam nisl, id suscipit tellus. Ut id tempor orci, non gravida odio.\n Nullam in scelerisque felis, id aliquet urna. Nam dapibus maximus ante, eu tempor neque lobortis eget.\n Suspendisse ut enim tincidunt, condimentum arcu id, ultrices urna. Maecenas sollicitudin fringilla aliquet.", gofpdf.BorderFull, gofpdf.AlignRight, false)

	/* ::Cosas de formas basicas:: */
	pdf.SetFillColor(0, 255, 0) //F
	pdf.SetDrawColor(0, 0, 255) //D
	//-> Dibujar rectangulo:
	pdf.Rect(10, 100, 100, 100, "FD") //<-(SetFillColor + SetDrawColor)

	pdf.SetFillColor(100, 200, 200) //F
	//-> Dibujar poligono:
	pdf.Polygon([]gofpdf.PointType{{110, 250}, {160, 300}, {110, 350}, {60, 300}}, "F")

	//-> Crear cuadricula:
	drawGrid(pdf)

	pdf.ImageOptions("imagenes/train.png", 275, 275, 92, 0, false, gofpdf.ImageOptions{ReadDpi: true}, 0, "")

	err := pdf.OutputFileAndClose("pf1.pdf")
	if err != nil {
		panic(err)
	}
}

func drawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
	for x := 0.0; x < w; x = x + (w / 20.0) { //del Letter size
		pdf.Line(x, 0, x, h)
		_, lineHeight := pdf.GetFontSize()
		pdf.Text(x, lineHeight, fmt.Sprintf("%d", int(x)))
	}

	for y := 0.0; y < h; y = y + (w / 20.0) { //del Letter size
		pdf.Line(0, y, w, y)
		//_, lineHeight := pdf.GetFontSize()
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}

}
