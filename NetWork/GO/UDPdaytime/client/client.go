package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

var service = ":1234"

func sendMsg(num int, wg *sync.WaitGroup) {

	// 1. use
	// udpAddr, err := net.ResolveUDPAddr("udp4", service)
	// checkError(err)
	// conn, err := net.DialUDP("udp", nil, udpAddr)
	// checkError(err)

	// 2. use Dial replace DialUDP
	conn, err := net.Dial("udp", service)
	checkError(err)
	msg := fmt.Sprintf("%d", num)
	_, err = conn.Write([]byte(msg))
	checkError(err)

	fmt.Printf("[*] Send %s\n", msg)

	var buf [512]byte
	n, err := conn.Read(buf[0:])

	checkError(err)
	fmt.Println(string(buf[0:n]))

	wg.Done()
}

func main() {
	var mux sync.WaitGroup
	for index := 0; index < 10; index++ {
		go sendMsg(index, &mux)
		mux.Add(1)
	}

	mux.Wait()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
