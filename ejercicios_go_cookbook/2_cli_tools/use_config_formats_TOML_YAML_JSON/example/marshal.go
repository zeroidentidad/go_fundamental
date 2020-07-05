package main

import (
	"configformats/json_"
	"configformats/toml_"
	"configformats/yaml_"
	"fmt"
)

// MarshalAll toma algunos datos almacenados en estructuras
// y los convierte a los distintos formatos de datos
func MarshalAll() error {
	t := toml_.TOMLData{
		Name: "Name1",
		Age:  20,
	}

	j := json_.JSONData{
		Name: "Name2",
		Age:  40,
	}

	y := yaml_.YAMLData{
		Name: "Name3",
		Age:  60,
	}

	tomlRes, err := t.ToTOML()
	if err != nil {
		return err
	}

	fmt.Println("TOML Marshal =", tomlRes.String())

	jsonRes, err := j.ToJSON()
	if err != nil {
		return err
	}

	fmt.Println("JSON Marshal=", jsonRes.String())

	yamlRes, err := y.ToYAML()
	if err != nil {
		return err
	}

	fmt.Println("YAML Marshal =", yamlRes.String())
	return nil
}
