package main

import (
	"fmt"
	"net"
	"os"
	"unicode/utf16"
)

const BOM = '\ufffe'

func main() {

	service := "127.0.0.1:1234"

	conn, err := net.Dial("tcp", service)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	shorts := readShorts(conn)
	ints := utf16.Decode(shorts)
	str := string(ints)
	fmt.Println(str)
	os.Exit(0)
}

func readShorts(conn net.Conn) []uint16 {
	buf := make([]byte, 512)
	n, err := conn.Read(buf[0:2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		m, err := conn.Read(buf[n:])
		if m == 0 || err != nil {
			break
		}
		n += m
	}

	shorts := make([]uint16, n/2)
	if buf[0] == 0xff && buf[1] == 0xfe {
		for i := 2; i < n; i += 2 {
			shorts[i/2] = uint16(buf[i])<<8 + uint16(buf[i+1])
		}
	} else if buf[0] == 0xfe && buf[1] == 0xff {
		for i := 2; i < n; i += 2 {
			shorts[i/2] = uint16(buf[i+1])<<8 + uint16(buf[i])
		}
	} else {
		fmt.Println("Unknown order")
	}
	return shorts
}
