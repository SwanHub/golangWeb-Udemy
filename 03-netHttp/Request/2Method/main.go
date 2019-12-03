package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type customStuff int

// retrieving the Method from the Request ("GET")

func (m customStuff) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		Submissions url.Values
	}{
		req.Method,
		req.Form,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var c customStuff
	http.ListenAndServe(":8080", c)
}
