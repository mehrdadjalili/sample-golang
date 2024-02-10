package grpc

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"os"
	"strings"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing-metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid-token")
)

func clientAuthentication(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	if !valid(md["authorization"], "client") {
		return nil, errInvalidToken
	}
	return handler(ctx, req)
}

func managerAuthentication(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	if !valid(md["authorization"], "manager") {
		return nil, errInvalidToken
	}
	return handler(ctx, req)
}

func valid(authorization []string, Type string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	key := getKey(Type)
	if key == "" {
		return false
	}
	return token == getKey(Type)
}

func getKey(Type string) string {
	if Type == "client" {
		return os.Getenv("CLIENT_AUTHENTICATION_KEY")
	} else if Type == "manager" {
		return os.Getenv("MANAGER_AUTHENTICATION_KEY")
	}
	return ""
}

func loadTLSCertificate() tls.Certificate {
	cert, _ := tls.LoadX509KeyPair(
		"resources/keys/tls/certificate.pem",
		"resources/keys/tls/key.pem",
	)
	return cert
}
