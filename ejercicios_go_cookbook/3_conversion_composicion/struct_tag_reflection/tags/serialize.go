package tags

import "reflect"

// SerializeStructStrings convierte una estructura
// al formato de serialización personalizado
// respeta las etiquetas de estructura de serialización
// para los tipos string
func SerializeStructStrings(s interface{}) (string, error) {
	result := ""

	// reflejar la interfaz en un tipo
	r := reflect.TypeOf(s)
	value := reflect.ValueOf(s)

	// si se pasa un puntero a una estructura, tratar adecuadamente
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
		value = value.Elem()
	}

	// recorrer todos los campos
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		// etiqueta de estructura encontrada
		key := field.Name
		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			// ignorar "-" de lo contrario todo ese valor
			// se convierte en la 'clave' de serialización
			if serialize == "-" {
				continue
			}
			key = serialize
		}

		switch value.Field(i).Kind() {
		// este ejemplo solo admite strings
		case reflect.String:
			result += key + ":" + value.Field(i).String() + ";"
			// por omisión
		default:
			continue
		}
	}
	return result, nil
}
