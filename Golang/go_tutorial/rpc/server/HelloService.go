package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const PingServiceName = "path/to/pkg.PingService"

type ServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterPingService(svc ServiceInterface) error {
	return rpc.RegisterName(PingServiceName, svc)
}

type HelloService struct{}

func (p *HelloService) Hello(request *String, reply *String) error {
	reply.Value = "server: " + request.GetValue()
	return nil
}

// func main() {
// 	RegisterPingService(new(PingService))

// 	listener, err := net.Listen("tcp", "localhost:1234")
// 	if err != nil {
// 		log.Fatal("ListenTCP error:", err)
// 	}

// 	for {
// 		conn, err := listener.Accept()
// 		fmt.Printf("[*]Get Request From: %s\n", conn.RemoteAddr().String())
// 		if err != nil {
// 			log.Fatal("Accept error:", err)
// 		}
// 		go rpc.ServeConn(conn)
// 	}
// }

func main() {
	rpc.RegisterName(PingServiceName, new(PingService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
