package templates

import (
	"fmt"
	"html/template"
	"os"
)

// HTMLDifferences destaca algunas de las diferencias
// entre html/template y text/template
func HTMLDifferences() error {
	t := template.New("html")
	t, err := t.Parse("<h1>Hello! {{.Name}}</h1>\n")
	if err != nil {
		return err
	}

	// html/template escapa automáticamente de operaciones inseguras como la inyección de JavaScript
	// esto es contextual y se comportará de manera diferente
	// dependiendo de dónde se representa una variable
	err = t.Execute(os.Stdout, map[string]string{"Name": "<script>alert('Can you see me?')</script>"})
	if err != nil {
		return err
	}

	// también puedes llamar manualmente a los escapers.
	fmt.Println(template.JSEscaper(`example <example@example.com>`))
	fmt.Println(template.HTMLEscaper(`example <example@example.com>`))
	fmt.Println(template.URLQueryEscaper(`example <example@example.com>`))

	return nil
}
