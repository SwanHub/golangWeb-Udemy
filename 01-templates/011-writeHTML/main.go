package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Templates
// to create (WRITE) an html page, run `go run main.go`
// to open the index page up in the browser, run `lite-server`

func main() {
	name := "Todd Withers"
	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Golang HTML test</title>
		</head>
		<body>
		<h1>I wrote this in Golang and my name is ` +
		name + `</h1>
		</body>
		</html>
		`)

	// create a document by name index.html
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	// defer closing the writer until after the function is over i.e. after writing is complete
	defer nf.Close()
	io.Copy(nf, strings.NewReader(str))

}
