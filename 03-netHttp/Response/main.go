package main

import (
	"fmt"
	"net/http"
)

// basic custom writer

type customGucci int

func (c customGucci) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Withers-Key", "this be from yb Tod")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Good old fashioned html through a go writer</h1>")
}

func main() {
	var gucci customGucci
	http.ListenAndServe(":8080", gucci)
}
