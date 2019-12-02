package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// Passing the data "Jimmy Gains" into the html variable $nicevar
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Jimmy Gains")
	if err != nil {
		log.Fatalln(err)
	}
}
