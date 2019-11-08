package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	address := ":1234"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		conn.SetDeadline(time.Now().Add(1000000000000))
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		fmt.Println(string(buf[0:]))
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
