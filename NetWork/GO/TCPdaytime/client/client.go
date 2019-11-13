package main

import (
	"fmt"
	"net"
	"os"
)

func getDate(num int) {
	service := "127.0.0.1:1234"

	// 1. use DailTCP
	// tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	// checkError(err)
	// conn, err := net.DialTCP("tcp", nil, tcpAddr)
	// checkError(err)

	// 2. use Dial replaces DailTCP
	conn, dialErr := net.Dial("tcp", service)
	checkError(dialErr)

	msg := fmt.Sprintf("%d", num)
	_, writeErr := conn.Write([]byte(msg))
	checkError(writeErr)
	fmt.Printf("[*] Send %s\n", msg)

	var buf [512]byte
	_, readErr := conn.Read(buf[0:])
	checkError(readErr)
	fmt.Printf("[*] Get %s\n", buf)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	getDate(1)
}
