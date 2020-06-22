package csvformat

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// Movie tomara CSV analizado
type Movie struct {
	Title    string
	Director string
	Year     int
}

// ReadCSV da muestra algunos ejemplos de procesamiento CSV
// que se pasa como io.Reader
func ReadCSV(b io.Reader) ([]Movie, error) {

	r := csv.NewReader(b)

	// Estas son algunas opciones de configuración opcionales.
	r.Comma = ';'
	r.Comment = '-'

	var movies []Movie

	// agarra e ignora el encabezado por ahora
	// también podemos usar esto para una clave de diccionario o alguna
	// otra forma de búsqueda
	_, err := r.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	// bucle hasta que esté todo procesado
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err
		}

		m := Movie{record[0], record[1], int(year)}
		movies = append(movies, m)
	}
	return movies, nil
}

// AddMoviesFromText usa el analizador CSV con una cadena
func AddMoviesFromText() error {
	// este es un ejemplo de nosotros tomando una cadena, convirtiendo
	// en un búfer y leyéndolo con el paquete csv
	in := `
- first our headers
movie title;director;year released
- then some data
Guardians of the Galaxy Vol. 2;James Gunn;2017
Star Wars: Episode VIII;Rian Johnson;2017
`

	b := bytes.NewBufferString(in)
	m, err := ReadCSV(b)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", m)
	return nil
}
