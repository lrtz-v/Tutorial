package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const PingServiceName = "path/to/pkg.PingService"

type PingServiceClient struct {
	*rpc.Client
}

// var _ PingServiceInterface = (*PingServiceClient)(nil)

func DialPingService(network, address string) (*PingServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &PingServiceClient{Client: c}, nil
}

func (p *PingServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(PingServiceName+".Hello", request, reply)
}

// func main() {
// 	client, err := DialPingService("tcp", "localhost:1234")
// 	if err != nil {
// 		log.Fatal("dialing:", err)
// 	}

// 	var reply string
// 	err = client.Hello("client", &reply)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("[*]Get Replay: %s\n", reply)
// }

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call(PingServiceName+".Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
