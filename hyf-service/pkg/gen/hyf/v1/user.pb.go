// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: hyf/v1/user.proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LastName string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_hyf_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetUsersIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUsersIn) Reset() {
	*x = GetUsersIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersIn) ProtoMessage() {}

func (x *GetUsersIn) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersIn.ProtoReflect.Descriptor instead.
func (*GetUsersIn) Descriptor() ([]byte, []int) {
	return file_hyf_v1_user_proto_rawDescGZIP(), []int{1}
}

type GetUsersOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetUsersOut) Reset() {
	*x = GetUsersOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersOut) ProtoMessage() {}

func (x *GetUsersOut) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersOut.ProtoReflect.Descriptor instead.
func (*GetUsersOut) Descriptor() ([]byte, []int) {
	return file_hyf_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUsersOut) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetUserIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserIn) Reset() {
	*x = GetUserIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserIn) ProtoMessage() {}

func (x *GetUserIn) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserIn.ProtoReflect.Descriptor instead.
func (*GetUserIn) Descriptor() ([]byte, []int) {
	return file_hyf_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserIn) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserOut) Reset() {
	*x = GetUserOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserOut) ProtoMessage() {}

func (x *GetUserOut) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserOut.ProtoReflect.Descriptor instead.
func (*GetUserOut) Descriptor() ([]byte, []int) {
	return file_hyf_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserOut) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateUserIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateUserIn) Reset() {
	*x = CreateUserIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserIn) ProtoMessage() {}

func (x *CreateUserIn) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserIn.ProtoReflect.Descriptor instead.
func (*CreateUserIn) Descriptor() ([]byte, []int) {
	return file_hyf_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserIn) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateUserOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateUserOut) Reset() {
	*x = CreateUserOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserOut) ProtoMessage() {}

func (x *CreateUserOut) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserOut.ProtoReflect.Descriptor instead.
func (*CreateUserOut) Descriptor() ([]byte, []int) {
	return file_hyf_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *CreateUserOut) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateUserIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateUserIn) Reset() {
	*x = UpdateUserIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserIn) ProtoMessage() {}

func (x *UpdateUserIn) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserIn.ProtoReflect.Descriptor instead.
func (*UpdateUserIn) Descriptor() ([]byte, []int) {
	return file_hyf_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateUserIn) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateUserOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateUserOut) Reset() {
	*x = UpdateUserOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserOut) ProtoMessage() {}

func (x *UpdateUserOut) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserOut.ProtoReflect.Descriptor instead.
func (*UpdateUserOut) Descriptor() ([]byte, []int) {
	return file_hyf_v1_user_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateUserOut) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type DeleteUserIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserIn) Reset() {
	*x = DeleteUserIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserIn) ProtoMessage() {}

func (x *DeleteUserIn) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserIn.ProtoReflect.Descriptor instead.
func (*DeleteUserIn) Descriptor() ([]byte, []int) {
	return file_hyf_v1_user_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteUserIn) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteUserOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserOut) Reset() {
	*x = DeleteUserOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserOut) ProtoMessage() {}

func (x *DeleteUserOut) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserOut.ProtoReflect.Descriptor instead.
func (*DeleteUserOut) Descriptor() ([]byte, []int) {
	return file_hyf_v1_user_proto_rawDescGZIP(), []int{10}
}

var File_hyf_v1_user_proto protoreflect.FileDescriptor

var file_hyf_v1_user_proto_rawDesc = []byte{
	0x0a, 0x11, 0x68, 0x79, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x79, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x22, 0x2a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x22, 0x1b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x27, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x12,
	0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2a, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x29, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2a, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x12, 0x19, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x1e, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x32, 0xe4, 0x03, 0x0a, 0x07, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x56, 0x43, 0x12, 0x68, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x1a, 0x0c,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f, 0x75, 0x74, 0x22, 0x41, 0x92, 0x41,
	0x2d, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x53, 0x56, 0x43, 0x12, 0x10, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x10, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x6a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4f, 0x75, 0x74, 0x22, 0x46, 0x92, 0x41, 0x30, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x53, 0x56,
	0x43, 0x12, 0x0e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65,
	0x72, 0x1a, 0x15, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65,
	0x72, 0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x12, 0x89, 0x01, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x1a, 0x0e, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x22, 0x5c, 0x92, 0x41, 0x45, 0x0a, 0x07,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x56, 0x43, 0x12, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x29, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x20, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x1a, 0x09, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x77, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x42, 0x92, 0x41,
	0x2c, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x53, 0x56, 0x43, 0x12, 0x0f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x10, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0d, 0x2a, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x2a,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_hyf_v1_user_proto_rawDescOnce sync.Once
	file_hyf_v1_user_proto_rawDescData = file_hyf_v1_user_proto_rawDesc
)

func file_hyf_v1_user_proto_rawDescGZIP() []byte {
	file_hyf_v1_user_proto_rawDescOnce.Do(func() {
		file_hyf_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_hyf_v1_user_proto_rawDescData)
	})
	return file_hyf_v1_user_proto_rawDescData
}

var file_hyf_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_hyf_v1_user_proto_goTypes = []any{
	(*User)(nil),          // 0: User
	(*GetUsersIn)(nil),    // 1: GetUsersIn
	(*GetUsersOut)(nil),   // 2: GetUsersOut
	(*GetUserIn)(nil),     // 3: GetUserIn
	(*GetUserOut)(nil),    // 4: GetUserOut
	(*CreateUserIn)(nil),  // 5: CreateUserIn
	(*CreateUserOut)(nil), // 6: CreateUserOut
	(*UpdateUserIn)(nil),  // 7: UpdateUserIn
	(*UpdateUserOut)(nil), // 8: UpdateUserOut
	(*DeleteUserIn)(nil),  // 9: DeleteUserIn
	(*DeleteUserOut)(nil), // 10: DeleteUserOut
	(*emptypb.Empty)(nil), // 11: google.protobuf.Empty
}
var file_hyf_v1_user_proto_depIdxs = []int32{
	0,  // 0: GetUsersOut.users:type_name -> User
	0,  // 1: GetUserOut.user:type_name -> User
	0,  // 2: CreateUserIn.user:type_name -> User
	0,  // 3: CreateUserOut.user:type_name -> User
	0,  // 4: UpdateUserIn.user:type_name -> User
	0,  // 5: UpdateUserOut.user:type_name -> User
	1,  // 6: UserSVC.GetUsers:input_type -> GetUsersIn
	3,  // 7: UserSVC.GetUser:input_type -> GetUserIn
	7,  // 8: UserSVC.UpdateUser:input_type -> UpdateUserIn
	9,  // 9: UserSVC.DeleteUser:input_type -> DeleteUserIn
	2,  // 10: UserSVC.GetUsers:output_type -> GetUsersOut
	4,  // 11: UserSVC.GetUser:output_type -> GetUserOut
	8,  // 12: UserSVC.UpdateUser:output_type -> UpdateUserOut
	11, // 13: UserSVC.DeleteUser:output_type -> google.protobuf.Empty
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_hyf_v1_user_proto_init() }
func file_hyf_v1_user_proto_init() {
	if File_hyf_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hyf_v1_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
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
		file_hyf_v1_user_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetUsersIn); i {
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
		file_hyf_v1_user_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetUsersOut); i {
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
		file_hyf_v1_user_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserIn); i {
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
		file_hyf_v1_user_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserOut); i {
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
		file_hyf_v1_user_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CreateUserIn); i {
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
		file_hyf_v1_user_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CreateUserOut); i {
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
		file_hyf_v1_user_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateUserIn); i {
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
		file_hyf_v1_user_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateUserOut); i {
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
		file_hyf_v1_user_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteUserIn); i {
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
		file_hyf_v1_user_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteUserOut); i {
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
			RawDescriptor: file_hyf_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hyf_v1_user_proto_goTypes,
		DependencyIndexes: file_hyf_v1_user_proto_depIdxs,
		MessageInfos:      file_hyf_v1_user_proto_msgTypes,
	}.Build()
	File_hyf_v1_user_proto = out.File
	file_hyf_v1_user_proto_rawDesc = nil
	file_hyf_v1_user_proto_goTypes = nil
	file_hyf_v1_user_proto_depIdxs = nil
}
