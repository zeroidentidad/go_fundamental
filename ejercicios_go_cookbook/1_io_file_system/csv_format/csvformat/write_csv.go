package csvformat

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

// Book tiene autor y título
type Book struct {
	Author string
	Title  string
}

// Books es un tipo nombrado para array de Book
type Books []Book

// ToCSV toma un conjunto de libros y escribe en un io.Writer
// devuelve cualquier error
func (books *Books) ToCSV(w io.Writer) error {
	n := csv.NewWriter(w)
	err := n.Write([]string{"Author", "Title"})
	if err != nil {
		return err
	}
	for _, book := range *books {
		err := n.Write([]string{book.Author, book.Title})
		if err != nil {
			return err
		}
	}

	n.Flush()
	return n.Error()
}

// WriteCSVOutput inicializa un conjunto de libros
// y escribe a os.Stdout
func WriteCSVOutput() error {
	b := Books{
		Book{
			Author: "F Scott Fitzgerald",
			Title:  "The Great Gatsby",
		},
		Book{
			Author: "J D Salinger",
			Title:  "The Catcher in the Rye",
		},
	}

	return b.ToCSV(os.Stdout)
}

// WriteCSVBuffer devuelve un buffer csv para
// un conjunto de libros
func WriteCSVBuffer() (*bytes.Buffer, error) {
	b := Books{
		Book{
			Author: "F Scott Fitzgerald",
			Title:  "The Great Gatsby",
		},
		Book{
			Author: "J D Salinger",
			Title:  "The Catcher in the Rye",
		},
	}

	w := &bytes.Buffer{}
	err := b.ToCSV(w)
	return w, err
}
