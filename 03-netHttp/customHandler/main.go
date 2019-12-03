package main

import "fmt"

import "net/http"

// type Handler interface is moi **importante**
// any type that has the signature ServeHTTP(ResponseWriter, *Request)

// here, we create our own handler, called customGoodies
type customGoodies int

func (c customGoodies) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Put anything you want in here... life is good homie.")
}

func main() {
	var c customGoodies
	http.ListenAndServe(":8080", c)
}
