package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	fmt.Println("Listening on port 8000")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("Error reading: %v", err)
			return
		}
		// fmt.Printf("Received : %s", string(buf[:n]))
		_, err = conn.Write(buf[:n])
		if err != nil {
			log.Println("Error writing:", err)
			return
		}

	}
}
