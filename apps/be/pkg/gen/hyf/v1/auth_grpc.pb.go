// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthSVC_Login_FullMethodName    = "/AuthSVC/Login"
	AuthSVC_Register_FullMethodName = "/AuthSVC/Register"
	AuthSVC_Refresh_FullMethodName  = "/AuthSVC/Refresh"
	AuthSVC_Logout_FullMethodName   = "/AuthSVC/Logout"
)

// AuthSVCClient is the client API for AuthSVC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthSVCClient interface {
	Login(ctx context.Context, in *LoginIn, opts ...grpc.CallOption) (*AuthOut, error)
	Register(ctx context.Context, in *RegisterIn, opts ...grpc.CallOption) (*AuthOut, error)
	Refresh(ctx context.Context, in *RefreshIn, opts ...grpc.CallOption) (*AuthOut, error)
	Logout(ctx context.Context, in *LogoutIn, opts ...grpc.CallOption) (*LogoutOut, error)
}

type authSVCClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthSVCClient(cc grpc.ClientConnInterface) AuthSVCClient {
	return &authSVCClient{cc}
}

func (c *authSVCClient) Login(ctx context.Context, in *LoginIn, opts ...grpc.CallOption) (*AuthOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthOut)
	err := c.cc.Invoke(ctx, AuthSVC_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authSVCClient) Register(ctx context.Context, in *RegisterIn, opts ...grpc.CallOption) (*AuthOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthOut)
	err := c.cc.Invoke(ctx, AuthSVC_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authSVCClient) Refresh(ctx context.Context, in *RefreshIn, opts ...grpc.CallOption) (*AuthOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthOut)
	err := c.cc.Invoke(ctx, AuthSVC_Refresh_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authSVCClient) Logout(ctx context.Context, in *LogoutIn, opts ...grpc.CallOption) (*LogoutOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogoutOut)
	err := c.cc.Invoke(ctx, AuthSVC_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthSVCServer is the server API for AuthSVC service.
// All implementations must embed UnimplementedAuthSVCServer
// for forward compatibility.
type AuthSVCServer interface {
	Login(context.Context, *LoginIn) (*AuthOut, error)
	Register(context.Context, *RegisterIn) (*AuthOut, error)
	Refresh(context.Context, *RefreshIn) (*AuthOut, error)
	Logout(context.Context, *LogoutIn) (*LogoutOut, error)
	mustEmbedUnimplementedAuthSVCServer()
}

// UnimplementedAuthSVCServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthSVCServer struct{}

func (UnimplementedAuthSVCServer) Login(context.Context, *LoginIn) (*AuthOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthSVCServer) Register(context.Context, *RegisterIn) (*AuthOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthSVCServer) Refresh(context.Context, *RefreshIn) (*AuthOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedAuthSVCServer) Logout(context.Context, *LogoutIn) (*LogoutOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAuthSVCServer) mustEmbedUnimplementedAuthSVCServer() {}
func (UnimplementedAuthSVCServer) testEmbeddedByValue()                 {}

// UnsafeAuthSVCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthSVCServer will
// result in compilation errors.
type UnsafeAuthSVCServer interface {
	mustEmbedUnimplementedAuthSVCServer()
}

func RegisterAuthSVCServer(s grpc.ServiceRegistrar, srv AuthSVCServer) {
	// If the following call pancis, it indicates UnimplementedAuthSVCServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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

func _AuthSVC_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthSVCServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthSVC_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthSVCServer).Refresh(ctx, req.(*RefreshIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthSVC_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthSVCServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthSVC_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthSVCServer).Logout(ctx, req.(*LogoutIn))
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
		{
			MethodName: "Refresh",
			Handler:    _AuthSVC_Refresh_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AuthSVC_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hyf/v1/auth.proto",
}
