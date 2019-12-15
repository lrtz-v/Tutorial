package main

import (
	"encoding/json"
	"fmt"
	"github.com/Tutorial/NetWork/GO/json_server/common"
	"net"
)

func main() {
	server := "127.0.0.1:1234"
	listener, err := net.Listen("tcp", server)
	common.ErrCheck(err)

	for {
		conn, err := listener.Accept()
		common.ErrCheck(err)

		go handler(conn)
	}
}

func handler(conn net.Conn) {
	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	var person common.Person
	err := decoder.Decode(&person)
	common.ErrCheck(err)
	fmt.Println(person)
	err = encoder.Encode(person)
	common.ErrCheck(err)
	conn.Close()
}
