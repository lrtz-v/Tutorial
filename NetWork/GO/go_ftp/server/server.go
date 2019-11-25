package main

import (
	"fmt"
	"net"
	"os"

	"github.com/Tutorial/NetWork/GO/go_ftp/common"
)

func main() {
	server := "127.0.0.1:1234"
	listener, err := net.Listen("tcp", server)
	common.ErrCheck(err)

	for {
		conn, err1 := listener.Accept()
		common.ErrCheck(err1)

		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 512)
	for {
		n, err := conn.Read(buf[0:])
		common.ErrCheck(err)

		s := string(buf[0:n])
		fmt.Println("[*] Get command: " + s)
		if s[0:2] == common.CD {
			chdir(conn, s[3:])
		} else if s[0:3] == common.DIR {
			dirList(conn)
		} else if s[0:3] == common.PWD {
			pwd(conn)
		}
	}
}

func chdir(conn net.Conn, s string) {
	if os.Chdir(s) == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("ERROR"))
	}
}

func pwd(conn net.Conn) {
	s, err := os.Getwd()
	common.ErrCheck(err)
	conn.Write([]byte(s))
}

func dirList(conn net.Conn) {
	defer conn.Write([]byte("\r\n"))

	dir, err := os.Open(".")
	common.ErrCheck(err)

	names, err1 := dir.Readdirnames(-1)
	common.ErrCheck(err1)
	fmt.Println(names)

	for _, nm := range names {
		conn.Write([]byte(nm + "\r\n"))
	}
}
