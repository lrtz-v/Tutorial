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
	conn, err := net.Dial("tcp", server)
	checkError(err)

	var buf [512]byte
	_, err = conn.Read(buf[0:])
	checkError(err)

	var newTime time.Time
	_, err = asn1.Unmarshal(buf[:], &newTime)
	checkError(err)
	fmt.Println(newTime)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func errCheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
