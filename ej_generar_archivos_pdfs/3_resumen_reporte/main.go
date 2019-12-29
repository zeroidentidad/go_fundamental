package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

const (
	bannerHt = 94.0
	xIndent  = 40.0
)

func main() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v, height=%v\n", w, h)
	pdf.AddPage()

	// Top encabezado con color
	pdf.SetFillColor(72, 16, 20)
	pdf.Polygon([]gofpdf.PointType{
		{0, 0},
		{w, 0},
		{w, bannerHt},
		{0, bannerHt * 0.93}, // menor mañano proporcional
	}, "F")
	pdf.Polygon([]gofpdf.PointType{
		{0, h},
		{0, h - (bannerHt * 0.2)},
		{w, h - (bannerHt * 0.1)},
		{w, h},
	}, "F")

	//Contenido de prueba para ubicar en posiciones con ayuda de la funcion de cuadricula

	// Banner - REPORTE
	pdf.SetFont("arial", "B", 40)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt := pdf.GetFontSize()
	pdf.Text(xIndent, bannerHt-(bannerHt/2.0)+lineHt/3.1, "REPORTE")

	// Banner - telefono, email, website
	pdf.SetFont("arial", "", 12)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt = pdf.GetFontSize()
	pdf.MoveTo(w-xIndent-2.0*124.0, (bannerHt-(lineHt*1.5*3.0))/2.0)
	pdf.MultiCell(124.0, lineHt*1.5, "(123) 456-7890\ntest@ejemplo.com\nejemplo.com", gofpdf.BorderNone, gofpdf.AlignRight, false)

	// Banner - direccion
	pdf.SetFont("arial", "", 12)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt = pdf.GetFontSize()
	pdf.MoveTo(w-xIndent-124.0, (bannerHt-(lineHt*1.5*3.0))/2.0)
	pdf.MultiCell(124.0, lineHt*1.5, "123 Lugar x\nAlguna ciudad, Calle\n12345", gofpdf.BorderNone, gofpdf.AlignRight, false)

	// Summary - Facturado a, número de factura, fecha de emisión
	_, sy := summaryBlock(pdf, xIndent, bannerHt+lineHt*2.0, "Facturado a", "Nombre cliente", "123 Direccion cliente", "Ciudad, Estado, Pais", "Codigo postal")
	summaryBlock(pdf, xIndent*2.0+lineHt*12.5, bannerHt+lineHt*2.0, "Numero factura", "0000000123")
	summaryBlock(pdf, xIndent*2.0+lineHt*12.5, bannerHt+lineHt*6.25, "Fecha de emision", "28/12/2019")

	// Summary - Total de la factura
	x, y := w-xIndent-124.0, bannerHt+lineHt*2.25
	pdf.MoveTo(x, y)
	pdf.SetFont("times", "", 14)
	_, lineHt = pdf.GetFontSize()
	pdf.SetTextColor(180, 180, 180)
	pdf.CellFormat(124.0, lineHt, "Total de factura", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")

	x, y = x+2.0, y+lineHt*1.5
	pdf.MoveTo(x, y)
	pdf.SetFont("times", "", 48)
	_, lineHt = pdf.GetFontSize()
	alpha := 58
	pdf.SetTextColor(72+alpha, 42+alpha, 55+alpha)
	totalUSD := "$1234.56"
	pdf.CellFormat(124.0, lineHt, totalUSD, gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x, y = x-2.0, y+lineHt*1.25

	if sy > y {
		y = sy
	}
	x, y = xIndent-20.0, y+30.0
	pdf.Rect(x, y, w-(xIndent*2.0)+40.0, 3.0, "F")

	// Grid util para las medidas de los elmentos y posicion
	// drawGrid(pdf)
	err := pdf.OutputFileAndClose("pf3.pdf")
	if err != nil {
		panic(err)
	}
}

func summaryBlock(pdf *gofpdf.Fpdf, x, y float64, title string, data ...string) (float64, float64) {
	pdf.SetFont("times", "", 14)
	pdf.SetTextColor(180, 180, 180)
	_, lineHt := pdf.GetFontSize()
	y = y + lineHt
	pdf.Text(x, y, title)
	y = y + lineHt*.25
	pdf.SetTextColor(50, 50, 50)
	for _, str := range data {
		y = y + lineHt*1.25
		pdf.Text(x, y, str)
	}
	return x, y
}

func drawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
	for x := 0.0; x < w; x = x + (w / 20.0) {
		pdf.SetTextColor(200, 200, 200)
		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}
	for y := 0.0; y < h; y = y + (w / 20.0) {
		if y < bannerHt*.9 {
			pdf.SetTextColor(200, 200, 200)
		} else {
			pdf.SetTextColor(80, 80, 80)
		}
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
}
