package main

import (
	"fmt"
	"net"
	"os"
)

func getDate() string {
	service := "127.0.0.1:1234"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)

	return string(buf[0:n])
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	getDate()
}
