package tags

import "fmt"

// Person es una estructura que almacena personas
// nombre, ciudad, estado y un atributo misceláneo
type Person struct {
	Name  string `serialize:"name"`
	City  string `serialize:"city"`
	State string
	Misc  string `serialize:"-"`
	Year  int    `serialize:"year"`
}

// EmptyStruct demuestra serializar
// y deserializar para una estructura vacía
// con etiquetas
func EmptyStruct() error {
	p := Person{}

	res, err := SerializeStructStrings(&p)
	if err != nil {
		return err
	}
	fmt.Printf("Empty struct: %#v\n", p)
	fmt.Println("Serialize results:", res)

	newP := Person{}
	if err := DeSerializeStructStrings(res, &newP); err != nil {
		return err
	}
	fmt.Printf("Deserialize results: %#v\n", newP)
	return nil
}

// FullStruct demuestra serializar
// y deserializar para una estructura completa
// con etiquetas
func FullStruct() error {
	p := Person{
		Name:  "Vicente",
		City:  "Cunduacan",
		State: "TAB",
		Misc:  "dejo este mundo",
		Year:  2020,
	}
	res, err := SerializeStructStrings(&p)
	if err != nil {
		return err
	}
	fmt.Printf("Full struct: %#v\n", p)
	fmt.Println("Serialize results:", res)

	newP := Person{}
	if err := DeSerializeStructStrings(res, &newP); err != nil {
		return err
	}
	fmt.Printf("Deserialize results: %#v\n", newP)
	return nil
}
