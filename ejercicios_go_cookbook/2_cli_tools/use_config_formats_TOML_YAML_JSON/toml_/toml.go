package toml_

import (
	"bytes"

	"github.com/BurntSushi/toml" //pgk externo más comun
)

// TOMLData es la estructura de datos común
// con etiquetas TOML struct
type TOMLData struct {
	Name string `toml:"name"`
	Age  int    `toml:"age"`
}

// ToTOML vuelca la estructura TOMLData a
// a bytes del formato TOML.
func (t *TOMLData) ToTOML() (*bytes.Buffer, error) {
	b := &bytes.Buffer{}
	encoder := toml.NewEncoder(b)

	if err := encoder.Encode(t); err != nil {
		return nil, err
	}
	return b, nil
}

// Decode se decodificará en TOMLData
func (t *TOMLData) Decode(data []byte) (toml.MetaData, error) {
	return toml.Decode(string(data), t)
}
