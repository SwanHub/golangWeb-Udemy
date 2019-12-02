package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// PARSED FILES
	// tpl == pointer to a template holding all of the files I've parsed
	tpl, err := template.ParseFiles("example.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// CREATE FILES
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	//EXECUTE FILES
	// because it's a pointer to a template, I can run the Execute function
	// which will write to the newly created file 'index.html'
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
