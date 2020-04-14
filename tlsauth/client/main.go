package main

import (
	"context"
	"fmt"
	"gomicroservice/tlsauth/message"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("../certs/server.pem", "CN")
	if err != nil {
		grpclog.Fatal("Load certs failed", err)
	}
	conn, err := grpc.Dial("127.0.0.1:8092", grpc.WithTransportCredentials(creds))
	if err != nil {
		grpclog.Fatal("Dial grpc server error", err)
	}
	defer conn.Close()

	serviceClient := message.NewMethServiceClient(conn)
	req := &message.RequestArgs{Arg1: 3.4, Arg2: 4.2}

	resp, err := serviceClient.AddMethod(context.Background(), req)
	if err != nil {
		grpclog.Fatal(err)
	}
	fmt.Printf("%+v", resp)
}
