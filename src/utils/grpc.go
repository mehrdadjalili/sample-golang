package utils

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"strings"
)

func GetGrpcAccessToken(ctx context.Context) (string, string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", "", errors.New("error")
	}
	authorization := md.Get("authorization")
	if len(authorization) < 1 {
		return "", "", errors.New("error")
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	tokens := strings.Split(token, ":")
	if len(tokens) < 2 {
		return "", "", errors.New("error")
	}
	return tokens[0], tokens[1], nil
}

func CreateGrpcDialOption(token, domain, certPath string) ([]grpc.DialOption, error) {
	certs := oauth.NewOauthAccess(&oauth2.Token{AccessToken: token})
	caCert, err := ioutil.ReadFile(certPath)
	if err != nil {
		return nil, err
	}
	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(caCert)
	if err != nil {
		return nil, err
	}
	tlsConf := &tls.Config{
		RootCAs:            rootCAs,
		InsecureSkipVerify: false,
		MinVersion:         tls.VersionTLS12,
		ServerName:         domain,
	}
	return []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConf)),
		grpc.WithPerRPCCredentials(certs),
	}, nil
}
