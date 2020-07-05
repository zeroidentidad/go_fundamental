package yaml_

import (
	"bytes"

	"gopkg.in/yaml.v2" //github.com/go-yaml/yaml
)

// YAMLData es la estructura de datos común
// con etiquetas de estruct YAML
type YAMLData struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

// ToYAML vuelca la estructura YAMLData a
// bytes de formato YAML.
func (t *YAMLData) ToYAML() (*bytes.Buffer, error) {
	d, err := yaml.Marshal(t)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(d)

	return b, nil
}

// Decode decodificará en TOMLData
func (t *YAMLData) Decode(data []byte) error {
	return yaml.Unmarshal(data, t)
}
