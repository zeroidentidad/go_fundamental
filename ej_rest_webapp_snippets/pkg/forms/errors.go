package forms

type formErrors map[string][]string

func (e formErrors) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e formErrors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}
