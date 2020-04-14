package main

import (
	gorder "gomicroservice/grpcrpc/order"
	"gomicroservice/grpcrpc/server/order"
	"net"

	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	gorder.RegisterOrderServiceServer(server, &order.OrderServiceImpl{})
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server.Serve(listen)
}
