package main

import (
	"encoding/json"
	"log"
	"net"
)

type request struct {
	Method string  `json:"method"`
	Number float64 `json:"number"`
}

type response struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}

func isPrime(number float64) bool {
	for i := 2; i < int(number); i++ {
		if int(number)%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	addr, err := net.ResolveTCPAddr("tcp", ":9080")
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

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

	var req request
	var res response

	for {
		decoder := json.NewDecoder(conn)
		err := decoder.Decode(&req)
		if err != nil {
			log.Fatal(err)
		}

		if req.Method != "isPrime" || req.Number == 0 {
			res = response{
				Method: "error",
				Prime:  false,
			}
			encoder := json.NewEncoder(conn)
			err = encoder.Encode(res)
			if err != nil {
				log.Println(err)
				continue
			}

			err = conn.Close()
		}

		res = response{
			Method: req.Method,
			Prime:  isPrime(req.Number),
		}

		encoder := json.NewEncoder(conn)
		err = encoder.Encode(res)
		if err != nil {
			log.Println(err)
		}
	}

}
