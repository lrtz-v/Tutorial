package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"github.com/Tutorial/NetWork/GO/gob_server/common"
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
	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	var person common.Person
	err := decoder.Decode(&person)
	common.ErrCheck(err)
	fmt.Println(person)
	err = encoder.Encode(person)
	common.ErrCheck(err)
	conn.Close()
}
