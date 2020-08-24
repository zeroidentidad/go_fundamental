package database

func Exec(db DB) error {

	// error no detectado en limpieza reejcución, pero siempre limpiar
	defer db.Exec("DROP TABLE example")

	if err := Create(db); err != nil {
		return err
	}

	if err := Query(db); err != nil {
		return err
	}
	return nil
}
