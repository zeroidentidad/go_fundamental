package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

type Form struct {
	url.Values
	Errors formErrors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		formErrors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)

		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "Este campo no puede estar vacío")
		}
	}
}

func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)

	if value == "" {
		return
	}

	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("Este campo es demasiado largo (máximo %d caracteres)", d))
	}
}

func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)

	if value == "" {
		return
	}

	for _, opt := range opts {
		if value == opt {
			return
		}
	}

	f.Errors.Add(field, "Este campo no es válido")
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
