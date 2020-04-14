package mytoken

import "context"

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
