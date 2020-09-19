package main

import (
	"crypto/tls"
	"fmt"
	"os"
)

func main() {
	service := "127.0.0.1:1234"

	conn, err := tls.Dial("tcp", service, nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for n := 0; n < 10; n++ {
		conn.Write([]byte("Hello " + string(n+48)))

		buf := make([]byte, 512)
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println(string(buf[0:n]))
	}
	os.Exit(0)
}
