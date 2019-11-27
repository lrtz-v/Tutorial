package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	cert, err := tls.LoadX509KeyPair("../security/github.name.pem", "../security/private.pem")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	config := tls.Config{Certificates: []tls.Certificate{cert}}

	now := time.Now()
	config.Time = func() time.Time { return now }
	config.Rand = rand.Reader

	service := "127.0.0.1:1234"

	listener, err1 := tls.Listen("tcp", service, &config)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(1)
	}
	for {
		conn, err3 := listener.Accept()
		if err3 != nil {
			fmt.Println(err3.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 512)
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}
