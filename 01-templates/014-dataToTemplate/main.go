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
	// Passing the data "Tod Withers" into the html at the '.' between {{}}
	// The equivalent of Ruby's ERB tags
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Tod Withers")
	if err != nil {
		log.Fatalln(err)
	}

}
