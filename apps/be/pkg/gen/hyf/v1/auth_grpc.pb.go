// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: hyf/v1/auth.proto

package proto

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
	AuthSVC_Login_FullMethodName    = "/AuthSVC/Login"
	AuthSVC_Register_FullMethodName = "/AuthSVC/Register"
)

// AuthSVCClient is the client API for AuthSVC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthSVCClient interface {
	Login(ctx context.Context, in *LoginIn, opts ...grpc.CallOption) (*AuthOut, error)
	Register(ctx context.Context, in *RegisterIn, opts ...grpc.CallOption) (*AuthOut, error)
}

type authSVCClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthSVCClient(cc grpc.ClientConnInterface) AuthSVCClient {
	return &authSVCClient{cc}
}

func (c *authSVCClient) Login(ctx context.Context, in *LoginIn, opts ...grpc.CallOption) (*AuthOut, error) {
	out := new(AuthOut)
	err := c.cc.Invoke(ctx, AuthSVC_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authSVCClient) Register(ctx context.Context, in *RegisterIn, opts ...grpc.CallOption) (*AuthOut, error) {
	out := new(AuthOut)
	err := c.cc.Invoke(ctx, AuthSVC_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthSVCServer is the server API for AuthSVC service.
// All implementations must embed UnimplementedAuthSVCServer
// for forward compatibility
type AuthSVCServer interface {
	Login(context.Context, *LoginIn) (*AuthOut, error)
	Register(context.Context, *RegisterIn) (*AuthOut, error)
	mustEmbedUnimplementedAuthSVCServer()
}

// UnimplementedAuthSVCServer must be embedded to have forward compatible implementations.
type UnimplementedAuthSVCServer struct {
}

func (UnimplementedAuthSVCServer) Login(context.Context, *LoginIn) (*AuthOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthSVCServer) Register(context.Context, *RegisterIn) (*AuthOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthSVCServer) mustEmbedUnimplementedAuthSVCServer() {}

// UnsafeAuthSVCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthSVCServer will
// result in compilation errors.
type UnsafeAuthSVCServer interface {
	mustEmbedUnimplementedAuthSVCServer()
}

func RegisterAuthSVCServer(s grpc.ServiceRegistrar, srv AuthSVCServer) {
	s.RegisterService(&AuthSVC_ServiceDesc, srv)
}

func _AuthSVC_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthSVCServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthSVC_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthSVCServer).Login(ctx, req.(*LoginIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthSVC_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthSVCServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthSVC_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthSVCServer).Register(ctx, req.(*RegisterIn))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthSVC_ServiceDesc is the grpc.ServiceDesc for AuthSVC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthSVC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuthSVC",
	HandlerType: (*AuthSVCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthSVC_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _AuthSVC_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hyf/v1/auth.proto",
}