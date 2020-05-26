package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "productinfo/client/ecommerce"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrderManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// order, err := c.GetOrder(ctx, &wrappers.StringValue{Value: "102"})
	// if err != nil {
	// 	log.Fatalf("Could not get order: %v", err)
	// }
	// log.Printf("Order: %v", order.String())

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

	channel := make(chan struct{})
	go asncClientBidirectionalRPC(streamProcOrder, channel)
	time.Sleep(time.Millisecond * 100)

	if err := streamProcOrder.Send(&wrappers.StringValue{Value: "101"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", c, "101", err)
	}

	if err := streamProcOrder.CloseSend(); err != nil {
		log.Fatal(err)
	}
	<-channel
}

func asncClientBidirectionalRPC(streamProcOrder pb.OrderManagement_ProcessOrdersClient, c chan struct{}) {
	for {
		combinedShipment, errProcOrder := streamProcOrder.Recv()
		if errProcOrder == io.EOF {
			break
		}
		log.Println("Combined shipment : ", combinedShipment.OrdersList)
	}
	<-c

}
