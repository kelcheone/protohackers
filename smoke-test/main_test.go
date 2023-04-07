package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"testing"
)

func TestEchoServer(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:9080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	fmt.Fprintf(conn, "Hello World from TCP client \n")

	// The server echoes the message back to the client
	// Read the response
	message, err := bufio.NewReader(conn).ReadString('\n')
	// message2, _ := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Message from server: " + message)
	// fmt.Print("Message from server: " + message2)
}
