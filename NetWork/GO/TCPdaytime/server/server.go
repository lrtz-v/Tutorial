package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	address := "127.0.0.1:1234"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Printf("[*] RemoteAddr: %s\n", conn.RemoteAddr().String())
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
