// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: hyf/v1/group.proto

package proto

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_hyf_v1_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetGroupsIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetGroupsIn) Reset() {
	*x = GetGroupsIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupsIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupsIn) ProtoMessage() {}

func (x *GetGroupsIn) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupsIn.ProtoReflect.Descriptor instead.
func (*GetGroupsIn) Descriptor() ([]byte, []int) {
	return file_hyf_v1_group_proto_rawDescGZIP(), []int{1}
}

type GetGroupsOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *GetGroupsOut) Reset() {
	*x = GetGroupsOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupsOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupsOut) ProtoMessage() {}

func (x *GetGroupsOut) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupsOut.ProtoReflect.Descriptor instead.
func (*GetGroupsOut) Descriptor() ([]byte, []int) {
	return file_hyf_v1_group_proto_rawDescGZIP(), []int{2}
}

func (x *GetGroupsOut) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

type GetGroupIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGroupIn) Reset() {
	*x = GetGroupIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupIn) ProtoMessage() {}

func (x *GetGroupIn) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupIn.ProtoReflect.Descriptor instead.
func (*GetGroupIn) Descriptor() ([]byte, []int) {
	return file_hyf_v1_group_proto_rawDescGZIP(), []int{3}
}

func (x *GetGroupIn) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetGroupOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *GetGroupOut) Reset() {
	*x = GetGroupOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupOut) ProtoMessage() {}

func (x *GetGroupOut) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupOut.ProtoReflect.Descriptor instead.
func (*GetGroupOut) Descriptor() ([]byte, []int) {
	return file_hyf_v1_group_proto_rawDescGZIP(), []int{4}
}

func (x *GetGroupOut) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type CreateGroupIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *CreateGroupIn) Reset() {
	*x = CreateGroupIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupIn) ProtoMessage() {}

func (x *CreateGroupIn) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupIn.ProtoReflect.Descriptor instead.
func (*CreateGroupIn) Descriptor() ([]byte, []int) {
	return file_hyf_v1_group_proto_rawDescGZIP(), []int{5}
}

func (x *CreateGroupIn) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type CreateGroupOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *CreateGroupOut) Reset() {
	*x = CreateGroupOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupOut) ProtoMessage() {}

func (x *CreateGroupOut) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupOut.ProtoReflect.Descriptor instead.
func (*CreateGroupOut) Descriptor() ([]byte, []int) {
	return file_hyf_v1_group_proto_rawDescGZIP(), []int{6}
}

func (x *CreateGroupOut) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type UpdateGroupIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *UpdateGroupIn) Reset() {
	*x = UpdateGroupIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupIn) ProtoMessage() {}

func (x *UpdateGroupIn) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupIn.ProtoReflect.Descriptor instead.
func (*UpdateGroupIn) Descriptor() ([]byte, []int) {
	return file_hyf_v1_group_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateGroupIn) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type UpdateGroupOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *UpdateGroupOut) Reset() {
	*x = UpdateGroupOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_group_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupOut) ProtoMessage() {}

func (x *UpdateGroupOut) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_group_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupOut.ProtoReflect.Descriptor instead.
func (*UpdateGroupOut) Descriptor() ([]byte, []int) {
	return file_hyf_v1_group_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateGroupOut) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type DeleteGroupIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGroupIn) Reset() {
	*x = DeleteGroupIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_group_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupIn) ProtoMessage() {}

func (x *DeleteGroupIn) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_group_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupIn.ProtoReflect.Descriptor instead.
func (*DeleteGroupIn) Descriptor() ([]byte, []int) {
	return file_hyf_v1_group_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteGroupIn) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteGroupOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteGroupOut) Reset() {
	*x = DeleteGroupOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_group_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupOut) ProtoMessage() {}

func (x *DeleteGroupOut) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_group_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupOut.ProtoReflect.Descriptor instead.
func (*DeleteGroupOut) Descriptor() ([]byte, []int) {
	return file_hyf_v1_group_proto_rawDescGZIP(), []int{10}
}

var File_hyf_v1_group_proto protoreflect.FileDescriptor

var file_hyf_v1_group_proto_rawDesc = []byte{
	0x0a, 0x12, 0x68, 0x79, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4d, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x0d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x49, 0x6e, 0x22, 0x2e,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x1e,
	0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x1c,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x2d, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x12, 0x1c, 0x0a, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x2e, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x2d, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x12, 0x1c, 0x0a, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x2e, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x75, 0x74, 0x32, 0x8c, 0x05, 0x0a, 0x08, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x53, 0x56, 0x43, 0x12, 0x6d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x12, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x49, 0x6e, 0x1a, 0x0d, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x4f, 0x75,
	0x74, 0x22, 0x43, 0x92, 0x41, 0x2e, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x11,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x1a, 0x11, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x6f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x1a,
	0x0c, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x75, 0x74, 0x22, 0x48, 0x92,
	0x41, 0x31, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x0f, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x73, 0x20, 0x61, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x16, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x73, 0x20, 0x61, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x20, 0x62, 0x79, 0x20,
	0x69, 0x64, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x2a, 0x12, 0x90, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x1a, 0x0f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x75, 0x74, 0x22, 0x60, 0x92, 0x41, 0x48, 0x0a, 0x06, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20,
	0x6e, 0x65, 0x77, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x2a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x20, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x20, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x8f, 0x01, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x1a, 0x0f, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x75, 0x74, 0x22, 0x5f, 0x92, 0x41, 0x47,
	0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x2b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x20, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a,
	0x1a, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x7b, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x44, 0x92, 0x41, 0x2d, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x12, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x1a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x2a, 0x0c, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x2a, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hyf_v1_group_proto_rawDescOnce sync.Once
	file_hyf_v1_group_proto_rawDescData = file_hyf_v1_group_proto_rawDesc
)

func file_hyf_v1_group_proto_rawDescGZIP() []byte {
	file_hyf_v1_group_proto_rawDescOnce.Do(func() {
		file_hyf_v1_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_hyf_v1_group_proto_rawDescData)
	})
	return file_hyf_v1_group_proto_rawDescData
}

var file_hyf_v1_group_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_hyf_v1_group_proto_goTypes = []interface{}{
	(*Group)(nil),          // 0: Group
	(*GetGroupsIn)(nil),    // 1: GetGroupsIn
	(*GetGroupsOut)(nil),   // 2: GetGroupsOut
	(*GetGroupIn)(nil),     // 3: GetGroupIn
	(*GetGroupOut)(nil),    // 4: GetGroupOut
	(*CreateGroupIn)(nil),  // 5: CreateGroupIn
	(*CreateGroupOut)(nil), // 6: CreateGroupOut
	(*UpdateGroupIn)(nil),  // 7: UpdateGroupIn
	(*UpdateGroupOut)(nil), // 8: UpdateGroupOut
	(*DeleteGroupIn)(nil),  // 9: DeleteGroupIn
	(*DeleteGroupOut)(nil), // 10: DeleteGroupOut
	(*emptypb.Empty)(nil),  // 11: google.protobuf.Empty
}
var file_hyf_v1_group_proto_depIdxs = []int32{
	0,  // 0: GetGroupsOut.groups:type_name -> Group
	0,  // 1: GetGroupOut.group:type_name -> Group
	0,  // 2: CreateGroupIn.group:type_name -> Group
	0,  // 3: CreateGroupOut.group:type_name -> Group
	0,  // 4: UpdateGroupIn.group:type_name -> Group
	0,  // 5: UpdateGroupOut.group:type_name -> Group
	1,  // 6: GroupSVC.GetGroups:input_type -> GetGroupsIn
	3,  // 7: GroupSVC.GetGroup:input_type -> GetGroupIn
	5,  // 8: GroupSVC.CreateGroup:input_type -> CreateGroupIn
	7,  // 9: GroupSVC.UpdateGroup:input_type -> UpdateGroupIn
	9,  // 10: GroupSVC.DeleteGroup:input_type -> DeleteGroupIn
	2,  // 11: GroupSVC.GetGroups:output_type -> GetGroupsOut
	4,  // 12: GroupSVC.GetGroup:output_type -> GetGroupOut
	6,  // 13: GroupSVC.CreateGroup:output_type -> CreateGroupOut
	8,  // 14: GroupSVC.UpdateGroup:output_type -> UpdateGroupOut
	11, // 15: GroupSVC.DeleteGroup:output_type -> google.protobuf.Empty
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_hyf_v1_group_proto_init() }
func file_hyf_v1_group_proto_init() {
	if File_hyf_v1_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hyf_v1_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hyf_v1_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupsIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hyf_v1_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupsOut); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hyf_v1_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hyf_v1_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupOut); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hyf_v1_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hyf_v1_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupOut); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hyf_v1_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hyf_v1_group_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupOut); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hyf_v1_group_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hyf_v1_group_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupOut); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hyf_v1_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hyf_v1_group_proto_goTypes,
		DependencyIndexes: file_hyf_v1_group_proto_depIdxs,
		MessageInfos:      file_hyf_v1_group_proto_msgTypes,
	}.Build()
	File_hyf_v1_group_proto = out.File
	file_hyf_v1_group_proto_rawDesc = nil
	file_hyf_v1_group_proto_goTypes = nil
	file_hyf_v1_group_proto_depIdxs = nil
}
