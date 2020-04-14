package mytoken

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type TokenAuthentication struct {
	AppKey    string
	AppSecret string
}

func (t *TokenAuthentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appkey":    t.AppKey,
		"appsecret": t.AppSecret,
	}, nil
}

func (t *TokenAuthentication) RequireTransportSecurity() bool {
	return true
}

func TokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
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

	return handler(ctx, req)
}
