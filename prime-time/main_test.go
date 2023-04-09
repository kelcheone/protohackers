package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestIsPrime(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:9080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	requests := []struct {
		method string
		number float64
	}{
		{"isPrime", 5},
		{"isPrime", 6},
		{"isPrime", 7},
		{"isPrime", 8},
		{"isPrime", 9},
		{"isPrime", 10},
		{"isPrime", 11},
		{"isPrime", 12},
	}

	expected := []struct {
		method string
		prime  bool
	}{
		{"isPrime", true},
		{"isPrime", false},
		{"isPrime", true},
		{"isPrime", false},
		{"isPrime", false},
		{"isPrime", false},
		{"isPrime", true},
		{"isPrime", false},
	}

	type response struct {
		Method string `json:"method"`
		Prime  bool   `json:"prime"`
	}

	for index, req := range requests {
		fmt.Fprintf(conn, `{"method":"%s","number":%f}`, req.method, req.number)
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		// convert JSON to struct
		var res response
		err := json.Unmarshal([]byte(msg), &res)
		if err != nil {
			log.Fatal(err)
		}
		if res.Method != expected[index].method || res.Prime != expected[index].prime {
			t.Errorf("Expected %v, got %v", expected[index], res)
			continue
		}

		assert.Equal(t, res.Method, expected[index].method)
		assert.Equal(t, res.Prime, expected[index].prime)

	}

}
