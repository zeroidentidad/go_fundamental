package repositories

import (
	"database/sql"
	"log"

	"../models"
)

//RepoLibro ...
type RepoLibro struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//GetLibros exported
func (rp RepoLibro) GetLibros(db *sql.DB, libro models.Libro, libros []models.Libro) []models.Libro {
	rows, err := db.Query("select * from libros")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&libro.ID, &libro.Titulo, &libro.Autor, &libro.Anio)
		logFatal(err)

		libros = append(libros, libro)
	}

	return libros
}

//GetLibro exported
func (rp RepoLibro) GetLibro(db *sql.DB, libro models.Libro, id int) models.Libro {
	rows := db.QueryRow("select * from libros where id=$1", id)

	err := rows.Scan(&libro.ID, &libro.Titulo, &libro.Autor, &libro.Anio)
	logFatal(err)

	return libro
}

//AddLibro exported
func (rp RepoLibro) AddLibro(db *sql.DB, libro models.Libro) int {
	err := db.QueryRow("insert into libros (titulo, autor, anio) values ($1, $2, $3) returning id;", libro.Titulo, libro.Autor, libro.Anio).Scan(&libro.ID)

	logFatal(err)

	return libro.ID
}

//UpdateLibro exported
func (rp RepoLibro) UpdateLibro(db *sql.DB, libro models.Libro) int64 {
	result, err := db.Exec("update libros set titulo=$1, autor=$2, anio=$3 where id=$4 returning id", libro.Titulo, libro.Autor, libro.Anio, libro.ID)
	logFatal(err)

	rowUpdated, err := result.RowsAffected()
	logFatal(err)

	return rowUpdated
}

//RemoveLibro exported
func (rp RepoLibro) RemoveLibro(db *sql.DB, id int) int64 {
	result, err := db.Exec("delete from libros where id=$1", id)
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	return rowsDeleted
}
