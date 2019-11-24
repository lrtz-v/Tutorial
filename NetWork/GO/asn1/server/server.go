package main

import (
	"encoding/asn1"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	server := "127.0.0.1:1234"
	listener, err := net.Listen("tcp", server)
	checkError(err)

	for {
		conn, err1 := listener.Accept()
		checkError(err1)

		go handler(conn)
	}
}

func handler(conn net.Conn) {
	dayTime := time.Now()
	mData, _ := asn1.Marshal(dayTime)
	conn.Write(mData)
	conn.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
