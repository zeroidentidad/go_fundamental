package tags

import (
	"errors"
	"reflect"
	"strings"
)

// DeSerializeStructStrings convierte una cadena serializada
// utilizando nuestro formato de serialización personalizado
// en una estructura
func DeSerializeStructStrings(s string, res interface{}) error {
	r := reflect.TypeOf(res)

	// configurando el uso de un puntero, por lo que
	// siempre debe ser un puntero pasado
	if r.Kind() != reflect.Ptr {
		return errors.New("res debe ser un puntero")
	}

	// desreferenciar el puntero
	r = r.Elem()
	value := reflect.ValueOf(res).Elem()

	// dividir la cadena de serialización en
	// un mapa
	vals := strings.Split(s, ";")
	valMap := make(map[string]string)
	for _, v := range vals {
		keyval := strings.Split(v, ":")
		if len(keyval) != 2 {
			continue
		}
		valMap[keyval[0]] = keyval[1]
	}

	// iterar sobre los campos
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)

		// comprobar si en el conjunto de serialización
		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			// ignorar "-" de lo contrario todo ese valor
			// se convierte en la 'clave' de serialización
			if serialize == "-" {
				continue
			}
			// esta en el mapa
			if val, ok := valMap[serialize]; ok {
				value.Field(i).SetString(val)
			}
		} else if val, ok := valMap[field.Name]; ok {
			// está el nombre de campo en el mapa?
			value.Field(i).SetString(val)
		}
	}
	return nil
}
