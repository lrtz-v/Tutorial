package main

import (
	"context"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"

	hello_pb "google.golang.org/grpc/examples/helloworld/helloworld"
	pb "productinfo/client/ecommerce"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(orderUnaryClientInterceptor),
		grpc.WithStreamInterceptor(clientStreamInterceptor))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrderManagementClient(conn)
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()

	clientDeadline := time.Now().Add(time.Duration(5 * time.Second))
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	order, getErr := c.GetOrder(ctx, &wrappers.StringValue{Value: "-1"})
	if getErr != nil {
		errorCode := status.Code(getErr)
		if errorCode == codes.InvalidArgument {
			log.Printf("Invalid Argument Error: %s", errorCode)
			errorStatus := status.Convert(getErr)
			for _, d := range errorStatus.Details() {
				switch info := d.(type) {
				case *epb.BadRequest_FieldViolation:
					log.Printf("Request Field Invalid: %s", info)
				default:
					log.Printf("Unexpected error type: %s", info)
				}
			}
		} else {
			log.Printf("Unhandled error: %s", errorCode)
		}
	} else {
		log.Printf("Order: %v", order.String())
	}

	helloClient := hello_pb.NewGreeterClient(conn)
	_, err = helloClient.SayHello(ctx, &hello_pb.HelloRequest{Name: "gRPC up and Running"})
	if err != nil {
		log.Printf("[*] SayHello Error: %s", err)
	}
	// searchStream, _ := c.SearchOrders(ctx, &wrappers.StringValue{Value: "Google"})
	// for {
	// 	searchOrder, err := searchStream.Recv()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	log.Print("Search Result: ", searchOrder)
	// }

	// updOrder1 := pb.Order{Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Destination: "Mountain View, CA", Price: 2222.00}
	// updOrder2 := pb.Order{Id: "103", Items: []string{"Apple Watch S4"}, Destination: "San Jose, CA", Price: 3333.00}
	// updOrder3 := pb.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub"}, Destination: "Mountain View, CA", Price: 4444.00}

	// updateStream, err := c.UpdateOrders(ctx)
	// if err != nil {
	// 	log.Fatalf("%v.UpdateOrders(_) = _, %v", c, err)
	// }
	// // update order 1
	// if err := updateStream.Send(&updOrder1); err != nil {
	// 	log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder1, err)
	// }
	// // update order 2
	// if err := updateStream.Send(&updOrder2); err != nil {
	// 	log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder2, err)
	// }
	// // update order 3
	// if err := updateStream.Send(&updOrder3); err != nil {
	// 	log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder3, err)
	// }

	// updateRes, err := updateStream.CloseAndRecv()
	// if err != nil {
	// 	log.Fatalf("%v.CloseAndRecv() got error %v, want %v", updateStream, err, nil)
	// }
	// log.Printf("Update Orders Res: %s", updateRes)

	// order, err = c.GetOrder(ctx, &wrappers.StringValue{Value: "102"})
	// if err != nil {
	// 	log.Fatalf("Could not get order: %v", err)
	// }
	// log.Printf("Order: %v", order.String())

	// BI
	streamProcOrder, err := c.ProcessOrders(ctx)
	if err != nil {
		log.Fatalf("%v.ProcessOrders(_) = _, %v", c, err)
	}

	if err := streamProcOrder.Send(&wrappers.StringValue{Value: "102"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", c, "102", err)
	}

	if err := streamProcOrder.Send(&wrappers.StringValue{Value: "103"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", c, "103", err)
	}

	if err := streamProcOrder.Send(&wrappers.StringValue{Value: "104"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", c, "104", err)
	}

	channel := make(chan int)
	go asncClientBidirectionalRPC(streamProcOrder, channel)
	time.Sleep(time.Millisecond * 1000)

	// Canceling the RPC
	cancel()
	log.Printf("RPC Status : %s", ctx.Err())
	if err := streamProcOrder.Send(&wrappers.StringValue{Value: "101"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", c, "101", err)
	}

	if err := streamProcOrder.CloseSend(); err != nil {
		log.Fatal(err)
	}
	<-channel
}

func asncClientBidirectionalRPC(streamProcOrder pb.OrderManagement_ProcessOrdersClient, c chan int) {
	for {
		combinedShipment, errProcOrder := streamProcOrder.Recv()
		if errProcOrder == io.EOF {
			break
		}
		log.Println("Combined shipment : ", combinedShipment.OrdersList)
	}
	c <- 0

}
func orderUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// Pre-processor phase
	log.Println("Method : " + method)

	// Invoking the remote method
	err := invoker(ctx, method, req, reply, cc, opts...)

	// Post-processor phase
	log.Println(reply)

	return err
}

func clientStreamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {

	log.Println("======= [Client Interceptor] ", method)
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrappedStream(s), nil
}

type wrappedStream struct {
	grpc.ClientStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("====== [Client Stream Interceptor] Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("====== [Client Stream Interceptor] Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}
