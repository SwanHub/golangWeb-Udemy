package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

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
		go handle(conn)
	}
}

// I was unable to verify this code as I can't use "telnet" in my cmd
// I see how it works, though, and am moving on
// I'm going to move forward speedily through TCP server action

// note that a mux is a router is a server is a multiplexer...

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		// write to the connection
		fmt.Fprintf(conn, "I heard you say %s\n", line)
	}
	defer conn.Close()
}
