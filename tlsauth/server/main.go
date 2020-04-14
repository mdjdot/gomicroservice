package main

import (
	"context"
	"gomicroservice/tlsauth/message"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

type MethManager struct {
}

func (m *MethManager) AddMethod(ctx context.Context, request *message.RequestArgs) (responste *message.ResponseResult, err error) {
	result := request.Arg1 + request.Arg2
	reponse := &message.ResponseResult{
		Result: result,
		Code:   1,
		Msg:    "ok",
	}
	return reponse, nil
}

func main() {
	creds, err := credentials.NewServerTLSFromFile("../certs/server.pem", "../certs/server.key")
	if err != nil {
		grpclog.Fatal("Load cert failed", err)
	}

	server := grpc.NewServer(grpc.Creds(creds))
	message.RegisterMethServiceServer(server, &MethManager{})

	listen, err := net.Listen("tcp", "127.0.0.1:8092")
	if err != nil {
		grpclog.Fatal("Listen port error", err)
	}
	server.Serve(listen)
}
