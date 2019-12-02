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
	// Passing data from anonymous struct into template
	info := []struct {
		First string
		Last  string
	}{
		{"tod", "withers"},
		{"jen", "poochinski"},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", info)
	if err != nil {
		log.Fatalln(err)
	}
}
