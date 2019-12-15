package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/Tutorial/NetWork/GO/go_ftp/common"
)

func main() {
	server := "127.0.0.1:1234"
	conn, err := net.Dial("tcp", server)
	common.ErrCheck(err)

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimRight(line, "\t\r\n")
		common.ErrCheck(err)

		strs := strings.SplitN(line, " ", 2)
		switch strs[0] {
		case common.DIR:
			dirRequest(conn)
		case common.CD:
			if len(strs) != 2 {
				fmt.Println("cd <dir>")
				continue
			}
			cdRequest(conn, strs[1])
		case common.PWD:
			pwdRequest(conn)
		case "quit":
			conn.Close()
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
		}
	}
}

func dirRequest(conn net.Conn) {
	conn.Write([]byte(common.DIR + " "))

	buf := make([]byte, 512)
	result := bytes.NewBuffer(nil)
	for {
		n, _ := conn.Read(buf[0:])
		result.Write(buf[0:n])
		length := result.Len()
		contents := result.Bytes()
		if string(contents[length-4:]) == "\r\n\r\n" {
			fmt.Println(string(contents[0 : length-4]))
			return
		}
	}
}

func cdRequest(conn net.Conn, dir string) {
	conn.Write([]byte(common.CD + " " + dir))
	response := make([]byte, 512)

	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	if s != "OK" {
		fmt.Println("Failed to change dir!")
	}
}

func pwdRequest(conn net.Conn) {
	conn.Write([]byte(common.PWD))
	response := make([]byte, 512)
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	fmt.Println("Current ddir \"" + s + "\"")
}
