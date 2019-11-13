package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

var address = "127.0.0.1:1234"

func main() {

	setupServer2()
}

func setupServer() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		go handler(conn)
	}
}

func handler(conn *net.TCPConn) {
	var buf [512]byte
	fmt.Printf("[*] RemoteAddr: %s\n", conn.RemoteAddr().String())
	_, err := conn.Read(buf[0:])
	if err != nil {
		return
	}
	fmt.Printf("[*] Receive %s\n", buf)
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
	conn.Close()
}

func setupServer2() {
	listener, err := net.Listen("tcp", address)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handler2(conn)
	}
}

func handler2(conn net.Conn) {
	var buf [512]byte
	fmt.Printf("[*] RemoteAddr: %s\n", conn.RemoteAddr().String())
	_, err := conn.Read(buf[0:])
	if err != nil {
		return
	}
	fmt.Printf("[*] Receive %s\n", buf)
	daytime := time.Now().String()
	returnData := []byte(daytime)
	conn.Write(returnData)
	fmt.Printf("[*] Return %s\n", returnData)
	conn.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
