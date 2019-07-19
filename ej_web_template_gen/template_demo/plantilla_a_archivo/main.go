package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
	"text/template"
)

type curso struct {
	Numero, Nombre, Unidades string
}

type semestre struct {
	Term   string
	Cursos []curso
}

func main() {
	// #1 get datos
	semestres := academicYear("datos/primer_semestre.txt", "datos/segundo_semestre.txt")

	// #2 parse template
	tpl, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// #3 create file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}

	// #4 execute template
	err = tpl.Execute(nf, semestres)
	if err != nil {
		log.Fatalln(err)
	}

}

func academicYear(s ...string) []semestre {
	semestres := []semestre{}
	for _, v := range s {
		s := semestre{}
		records := getRecords(v)
		xc := make([]curso, 0, len(records))
		// #3 get datos
		for i, row := range records {
			if i == 0 {
				xs := strings.SplitN(row[0], ",", 2)
				t := xs[0]
				s.Term = t
			} else {
				c := curso{}
				xs := strings.SplitN(row[0], " ", 2)
				c.Numero = xs[0]
				c.Nombre = xs[1]
				c.Unidades = row[1]
				xc = append(xc, c)
			}
		}
		s.Cursos = xc
		semestres = append(semestres, s)
	}
	return semestres
}

func getRecords(path string) [][]string {
	// #1 open a file
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// #2 parse a csv file
	rdr := csv.NewReader(f)
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}
	return rows
}
