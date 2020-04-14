package main

import (
	"context"
	"fmt"
	"gomicroservice/grpcrpc/order"
	"io"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	orderServiceClient := order.NewOrderServiceClient(conn)

	req := &order.OrderRequest{
		OrderID: "123",
	}
	resp, err := orderServiceClient.GetOrderInfo(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)

	client, err := orderServiceClient.GetOrderInfoStream(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		orderInfo, err := client.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", orderInfo)
	}

	clientC, err := orderServiceClient.GetOrderInfoClientStream(context.Background())
	if err != nil {
		panic(err)
	}

	reqMap := []order.OrderRequest{
		order.OrderRequest{OrderID: "1"},
		order.OrderRequest{OrderID: "2"},
		order.OrderRequest{OrderID: "3"},
	}
	for _, req := range reqMap {
		err = clientC.Send(&req)
		if err != nil {
			panic(err)
		}
	}

	for {
		orderInfo, err := clientC.CloseAndRecv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v\n", orderInfo)
	}

	client2, err := orderServiceClient.GetOrderInfo2Stream(context.Background())
	if err != nil {
		panic(err)
	}
	reqs := []string{"1", "2", "3"}
	for _, req := range reqs {
		request := order.OrderRequest{OrderID: req}
		err := client2.Send(&request)
		if err != nil {
			panic(err)
		}
	}
	client2.CloseSend()

	for {
		orderInfo, err := client2.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		fmt.Printf("%+v\n", orderInfo)
	}
}
