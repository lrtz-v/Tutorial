package main

import (
	"encoding/json"
	"fmt"
	"github.com/Tutorial/NetWork/GO/json/common"
	"net"
)

func main() {
	server := "127.0.0.1:1234"
	listener, err := net.Listen("tcp", server)
	ErrCheck(err)

	for {
		conn, err1 := listener.Accept()
		ErrCheck(err1)

		go handler(conn)
	}
}

func handler(conn net.Conn) {
	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	var person Person
	decoder.Decode(&person)
	fmt.Println(person)
	encoder.Encode(person)
	conn.Close()
}
