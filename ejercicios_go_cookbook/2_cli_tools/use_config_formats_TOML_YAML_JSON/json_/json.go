package json_

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JSONData es la estructura de datos común
// con etiquetas de estruct JSON
type JSONData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ToJSON vuelca la estructura JSONData a
// a bytes en formato JSON.
func (t *JSONData) ToJSON() (*bytes.Buffer, error) {
	d, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(d)

	return b, nil
}

// Decode decodificará en JSONData
func (t *JSONData) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}

// ================EXTRA====================
// OtherJSONExamples muestra formas de usar tipos
// más allá de estructuras y otras funciones útiles
func OtherJSONExamples() error {
	res := make(map[string]string)
	err := json.Unmarshal([]byte(`{"key": "value"}`), &res)
	if err != nil {
		return err
	}

	fmt.Println("Podemos descomponer en un mapa en lugar de una estructura:", res)

	b := bytes.NewReader([]byte(`{"key2": "value2"}`))
	decoder := json.NewDecoder(b)

	if err := decoder.Decode(&res); err != nil {
		return err
	}

	fmt.Println("También podemos usar decodificadores/codificadores para trabajar con streams:", res)

	return nil
}
