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
	// Passing data from slice into template
	homies := []string{"tod", "tanya", "gen", "poochinski"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", homies)
	if err != nil {
		log.Fatalln(err)
	}
}
