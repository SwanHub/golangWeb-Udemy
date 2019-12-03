package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// tcp is the transfer protocol for moving data back and forth
// http runs using tcp

// before using http, we're going to construct a server from the tcp up

func main() {
	// returns a Listener, which is an interface with methods "Accept" "Close" and "Addr"
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		// returns a Connection, which has Read and Write types
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nTCP Server running\n")
		fmt.Println("What's up?")
		fmt.Fprintf(conn, "%v", "Nothin")

		conn.Close()
	}
}
