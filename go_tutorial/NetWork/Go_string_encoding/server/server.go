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

	app, err := net.Listen("tcp", service)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		conn, err := app.Accept()
		if err != nil {
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	str := "おはようございます。"
	shorts := utf16.Encode([]rune(str))
	writeShorts(conn, shorts)
}

func writeShorts(conn net.Conn, shorts []uint16) {
	buf := make([]byte, 2)
	buf[0] = BOM >> 8
	buf[1] = BOM & 255
	_, err := conn.Write(buf)
	if err != nil {
		return
	}

	for _, v := range shorts {
		buf[0] = byte(v >> 8)
		buf[1] = byte(v & 255)
		_, err := conn.Write(buf)
		if err != nil {
			return
		}
	}
}
