package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func dayMonthYear(t time.Time) string {
	return t.Format("02-01-2006") // codigo de formato estandar -> 01-02-2006
}

var fm = template.FuncMap{
	"fdateDMY": dayMonthYear,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
