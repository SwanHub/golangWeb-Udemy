package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// three steps of hte server
// 1. Listen
// 2. Accept
// 3. Read and write to the connection

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		// using go routines to handle multiple connections
		// with each new li ping, we launch a new routine to handle the connection
		go handle(conn)
	}
}

// this function will print out the request SENT by the web browser to this server

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	// scanner.Scan() returns a bool... as long as there is text to read,
	// it will print the text and move to the next line
	for scanner.Scan() {
		// end of input, this will return false and the for loop will cease
		lineOfText := scanner.Text()
		fmt.Println(lineOfText)
	}

	// never gets past here... will always be listening... check 023 for how to remedy this
	defer conn.Close()
	fmt.Println("never gets here")
}
