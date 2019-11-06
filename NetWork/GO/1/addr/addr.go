package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", name)

	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved address is ", addr.String())

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}
	for i, s := range addrs {
		fmt.Printf("[*] %d: %s\n", i, s)
	}

	os.Exit(0)
}
