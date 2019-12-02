package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	// parse everything, and do it only once
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	// executes the first template in the template container that tpl points to
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// executes the file by the name of 'one.gohtml' to stdout
	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// executes the file by the name of 'two.gohtml' to stdout
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
