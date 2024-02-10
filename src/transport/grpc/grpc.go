package grpc

import (
	"fmt"
	"github.com/mehrdadjalili/facegram_auth_service/proto/pd_auth/pd_auth_manager"
	"github.com/mehrdadjalili/facegram_auth_service/src/service"
	"github.com/mehrdadjalili/facegram_auth_service/src/transport/grpc/pd_auth_client"
	"net"

	"github.com/mehrdadjalili/facegram_auth_service/src/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type (
	clientServer struct {
		grpc *grpc.Server
	}
	managerServer struct {
		grpc *grpc.Server
	}
)

func NewClient(authService service.Auth, userService service.User) transport.ClientGrpc {

	cert := loadTLSCertificate()

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.UnaryInterceptor(clientAuthentication),
	}

	s := grpc.NewServer(opts...)
	h := &clientHandler{
		authService: authService,
		userService: userService,
	}
	pd_auth_client.RegisterAuthClientServiceServer(s, h)
	return &clientServer{
		grpc: s,
	}
}

func NewManager(userService service.User, sessionService service.Session) transport.ManagerGrpc {

	cert := loadTLSCertificate()

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.UnaryInterceptor(managerAuthentication),
	}

	s := grpc.NewServer(opts...)
	h := &managerHandler{
		userService:    userService,
		sessionService: sessionService,
	}
	pd_auth_manager.RegisterAuthManagerServiceServer(s, h)
	return &managerServer{
		grpc: s,
	}
}

func (s *clientServer) StartClient(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	if err := s.grpc.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *managerServer) StartManager(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	if err := s.grpc.Serve(lis); err != nil {
		return err
	}
	return nil
}
