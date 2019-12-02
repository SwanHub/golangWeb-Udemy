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
	// Passing data from map into template
	info := map[string]string{
		"age":    "eleventy-one",
		"gender": "fluids",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", info)
	if err != nil {
		log.Fatalln(err)
	}
}
