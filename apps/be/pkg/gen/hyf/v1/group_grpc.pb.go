// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: hyf/v1/group.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GroupSVC_GetGroups_FullMethodName   = "/GroupSVC/GetGroups"
	GroupSVC_GetGroup_FullMethodName    = "/GroupSVC/GetGroup"
	GroupSVC_CreateGroup_FullMethodName = "/GroupSVC/CreateGroup"
	GroupSVC_UpdateGroup_FullMethodName = "/GroupSVC/UpdateGroup"
	GroupSVC_DeleteGroup_FullMethodName = "/GroupSVC/DeleteGroup"
)

// GroupSVCClient is the client API for GroupSVC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupSVCClient interface {
	GetGroups(ctx context.Context, in *GetGroupsIn, opts ...grpc.CallOption) (*GetGroupsOut, error)
	GetGroup(ctx context.Context, in *GetGroupIn, opts ...grpc.CallOption) (*GetGroupOut, error)
	CreateGroup(ctx context.Context, in *CreateGroupIn, opts ...grpc.CallOption) (*CreateGroupOut, error)
	UpdateGroup(ctx context.Context, in *UpdateGroupIn, opts ...grpc.CallOption) (*UpdateGroupOut, error)
	DeleteGroup(ctx context.Context, in *DeleteGroupIn, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type groupSVCClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupSVCClient(cc grpc.ClientConnInterface) GroupSVCClient {
	return &groupSVCClient{cc}
}

func (c *groupSVCClient) GetGroups(ctx context.Context, in *GetGroupsIn, opts ...grpc.CallOption) (*GetGroupsOut, error) {
	out := new(GetGroupsOut)
	err := c.cc.Invoke(ctx, GroupSVC_GetGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupSVCClient) GetGroup(ctx context.Context, in *GetGroupIn, opts ...grpc.CallOption) (*GetGroupOut, error) {
	out := new(GetGroupOut)
	err := c.cc.Invoke(ctx, GroupSVC_GetGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupSVCClient) CreateGroup(ctx context.Context, in *CreateGroupIn, opts ...grpc.CallOption) (*CreateGroupOut, error) {
	out := new(CreateGroupOut)
	err := c.cc.Invoke(ctx, GroupSVC_CreateGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupSVCClient) UpdateGroup(ctx context.Context, in *UpdateGroupIn, opts ...grpc.CallOption) (*UpdateGroupOut, error) {
	out := new(UpdateGroupOut)
	err := c.cc.Invoke(ctx, GroupSVC_UpdateGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupSVCClient) DeleteGroup(ctx context.Context, in *DeleteGroupIn, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GroupSVC_DeleteGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupSVCServer is the server API for GroupSVC service.
// All implementations must embed UnimplementedGroupSVCServer
// for forward compatibility
type GroupSVCServer interface {
	GetGroups(context.Context, *GetGroupsIn) (*GetGroupsOut, error)
	GetGroup(context.Context, *GetGroupIn) (*GetGroupOut, error)
	CreateGroup(context.Context, *CreateGroupIn) (*CreateGroupOut, error)
	UpdateGroup(context.Context, *UpdateGroupIn) (*UpdateGroupOut, error)
	DeleteGroup(context.Context, *DeleteGroupIn) (*emptypb.Empty, error)
	mustEmbedUnimplementedGroupSVCServer()
}

// UnimplementedGroupSVCServer must be embedded to have forward compatible implementations.
type UnimplementedGroupSVCServer struct {
}

func (UnimplementedGroupSVCServer) GetGroups(context.Context, *GetGroupsIn) (*GetGroupsOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (UnimplementedGroupSVCServer) GetGroup(context.Context, *GetGroupIn) (*GetGroupOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedGroupSVCServer) CreateGroup(context.Context, *CreateGroupIn) (*CreateGroupOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedGroupSVCServer) UpdateGroup(context.Context, *UpdateGroupIn) (*UpdateGroupOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedGroupSVCServer) DeleteGroup(context.Context, *DeleteGroupIn) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedGroupSVCServer) mustEmbedUnimplementedGroupSVCServer() {}

// UnsafeGroupSVCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupSVCServer will
// result in compilation errors.
type UnsafeGroupSVCServer interface {
	mustEmbedUnimplementedGroupSVCServer()
}

func RegisterGroupSVCServer(s grpc.ServiceRegistrar, srv GroupSVCServer) {
	s.RegisterService(&GroupSVC_ServiceDesc, srv)
}

func _GroupSVC_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupsIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupSVCServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupSVC_GetGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupSVCServer).GetGroups(ctx, req.(*GetGroupsIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupSVC_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupSVCServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupSVC_GetGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupSVCServer).GetGroup(ctx, req.(*GetGroupIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupSVC_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupSVCServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupSVC_CreateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupSVCServer).CreateGroup(ctx, req.(*CreateGroupIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupSVC_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupSVCServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupSVC_UpdateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupSVCServer).UpdateGroup(ctx, req.(*UpdateGroupIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupSVC_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupSVCServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupSVC_DeleteGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupSVCServer).DeleteGroup(ctx, req.(*DeleteGroupIn))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupSVC_ServiceDesc is the grpc.ServiceDesc for GroupSVC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupSVC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GroupSVC",
	HandlerType: (*GroupSVCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroups",
			Handler:    _GroupSVC_GetGroups_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _GroupSVC_GetGroup_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _GroupSVC_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _GroupSVC_UpdateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _GroupSVC_DeleteGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hyf/v1/group.proto",
}