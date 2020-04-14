package main

import (
	"context"
	"gomicroservice/tlsauth/message"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type MethManager struct {
}

func (m *MethManager) AddMethod(ctx context.Context, request *message.RequestArgs) (responste *message.ResponseResult, err error) {
	md, exit := metadata.FromIncomingContext(ctx)
	if !exit {
		return nil, status.Errorf(codes.Unauthenticated, "No token authentication into")
	}
	var appkey, appsecret string

	if key, ok := md["appkey"]; ok {
		appkey = key[0]
	}
	if secret, ok := md["appsecret"]; ok {
		appsecret = secret[0]
	}
	if appkey != "hello" || appsecret != "20190812" {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
	}

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
