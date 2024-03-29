// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: pd_auth_client.proto

package pd_auth_client

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AuthClientService_CheckToken_FullMethodName = "/pd_auth.AuthClientService/CheckToken"
	AuthClientService_UserById_FullMethodName   = "/pd_auth.AuthClientService/UserById"
)

// AuthClientServiceClient is the client API for AuthClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClientServiceClient interface {
	CheckToken(ctx context.Context, in *CheckTokenRequest, opts ...grpc.CallOption) (*CheckTokenResponse, error)
	UserById(ctx context.Context, in *ByUserIdRequest, opts ...grpc.CallOption) (*OneUserResponse, error)
}

type authClientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClientServiceClient(cc grpc.ClientConnInterface) AuthClientServiceClient {
	return &authClientServiceClient{cc}
}

func (c *authClientServiceClient) CheckToken(ctx context.Context, in *CheckTokenRequest, opts ...grpc.CallOption) (*CheckTokenResponse, error) {
	out := new(CheckTokenResponse)
	err := c.cc.Invoke(ctx, AuthClientService_CheckToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClientServiceClient) UserById(ctx context.Context, in *ByUserIdRequest, opts ...grpc.CallOption) (*OneUserResponse, error) {
	out := new(OneUserResponse)
	err := c.cc.Invoke(ctx, AuthClientService_UserById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthClientServiceServer is the server API for AuthClientService service.
// All implementations must embed UnimplementedAuthClientServiceServer
// for forward compatibility
type AuthClientServiceServer interface {
	CheckToken(context.Context, *CheckTokenRequest) (*CheckTokenResponse, error)
	UserById(context.Context, *ByUserIdRequest) (*OneUserResponse, error)
	mustEmbedUnimplementedAuthClientServiceServer()
}

// UnimplementedAuthClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthClientServiceServer struct {
}

func (UnimplementedAuthClientServiceServer) CheckToken(context.Context, *CheckTokenRequest) (*CheckTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckToken not implemented")
}
func (UnimplementedAuthClientServiceServer) UserById(context.Context, *ByUserIdRequest) (*OneUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserById not implemented")
}
func (UnimplementedAuthClientServiceServer) mustEmbedUnimplementedAuthClientServiceServer() {}

// UnsafeAuthClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthClientServiceServer will
// result in compilation errors.
type UnsafeAuthClientServiceServer interface {
	mustEmbedUnimplementedAuthClientServiceServer()
}

func RegisterAuthClientServiceServer(s grpc.ServiceRegistrar, srv AuthClientServiceServer) {
	s.RegisterService(&AuthClientService_ServiceDesc, srv)
}

func _AuthClientService_CheckToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthClientServiceServer).CheckToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthClientService_CheckToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthClientServiceServer).CheckToken(ctx, req.(*CheckTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthClientService_UserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthClientServiceServer).UserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthClientService_UserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthClientServiceServer).UserById(ctx, req.(*ByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthClientService_ServiceDesc is the grpc.ServiceDesc for AuthClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pd_auth.AuthClientService",
	HandlerType: (*AuthClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckToken",
			Handler:    _AuthClientService_CheckToken_Handler,
		},
		{
			MethodName: "UserById",
			Handler:    _AuthClientService_UserById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pd_auth_client.proto",
}
