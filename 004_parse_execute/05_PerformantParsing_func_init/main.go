package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func logIfError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	logIfError(tpl.Execute(os.Stdout, nil))

	logIfError(tpl.ExecuteTemplate(os.Stdout, "one.gahtml", nil))

	logIfError(tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil))

	logIfError(tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil))
}
