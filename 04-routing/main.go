package main

import "net/http"

import "io"

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "the d is silent")
}

func main() {
	http.HandleFunc("/boo", d)

	http.ListenAndServe(":8080", nil)
}
