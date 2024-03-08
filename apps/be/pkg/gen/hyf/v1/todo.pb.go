// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: hyf/v1/todo.proto

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

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_hyf_v1_todo_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Todo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Todo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetTodosIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTodosIn) Reset() {
	*x = GetTodosIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodosIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodosIn) ProtoMessage() {}

func (x *GetTodosIn) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodosIn.ProtoReflect.Descriptor instead.
func (*GetTodosIn) Descriptor() ([]byte, []int) {
	return file_hyf_v1_todo_proto_rawDescGZIP(), []int{1}
}

type GetTodosOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*Todo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *GetTodosOut) Reset() {
	*x = GetTodosOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodosOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodosOut) ProtoMessage() {}

func (x *GetTodosOut) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodosOut.ProtoReflect.Descriptor instead.
func (*GetTodosOut) Descriptor() ([]byte, []int) {
	return file_hyf_v1_todo_proto_rawDescGZIP(), []int{2}
}

func (x *GetTodosOut) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

type GetTodoIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTodoIn) Reset() {
	*x = GetTodoIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoIn) ProtoMessage() {}

func (x *GetTodoIn) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoIn.ProtoReflect.Descriptor instead.
func (*GetTodoIn) Descriptor() ([]byte, []int) {
	return file_hyf_v1_todo_proto_rawDescGZIP(), []int{3}
}

func (x *GetTodoIn) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTodoOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *GetTodoOut) Reset() {
	*x = GetTodoOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoOut) ProtoMessage() {}

func (x *GetTodoOut) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoOut.ProtoReflect.Descriptor instead.
func (*GetTodoOut) Descriptor() ([]byte, []int) {
	return file_hyf_v1_todo_proto_rawDescGZIP(), []int{4}
}

func (x *GetTodoOut) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type CreateTodoIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *CreateTodoIn) Reset() {
	*x = CreateTodoIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoIn) ProtoMessage() {}

func (x *CreateTodoIn) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoIn.ProtoReflect.Descriptor instead.
func (*CreateTodoIn) Descriptor() ([]byte, []int) {
	return file_hyf_v1_todo_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTodoIn) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type CreateTodoOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *CreateTodoOut) Reset() {
	*x = CreateTodoOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_todo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoOut) ProtoMessage() {}

func (x *CreateTodoOut) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_todo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoOut.ProtoReflect.Descriptor instead.
func (*CreateTodoOut) Descriptor() ([]byte, []int) {
	return file_hyf_v1_todo_proto_rawDescGZIP(), []int{6}
}

func (x *CreateTodoOut) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type UpdateTodoIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *UpdateTodoIn) Reset() {
	*x = UpdateTodoIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_todo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTodoIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoIn) ProtoMessage() {}

func (x *UpdateTodoIn) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_todo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoIn.ProtoReflect.Descriptor instead.
func (*UpdateTodoIn) Descriptor() ([]byte, []int) {
	return file_hyf_v1_todo_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTodoIn) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type UpdateTodoOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *UpdateTodoOut) Reset() {
	*x = UpdateTodoOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_todo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTodoOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoOut) ProtoMessage() {}

func (x *UpdateTodoOut) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_todo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoOut.ProtoReflect.Descriptor instead.
func (*UpdateTodoOut) Descriptor() ([]byte, []int) {
	return file_hyf_v1_todo_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateTodoOut) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type DeleteTodoIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTodoIn) Reset() {
	*x = DeleteTodoIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_todo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTodoIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodoIn) ProtoMessage() {}

func (x *DeleteTodoIn) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_todo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodoIn.ProtoReflect.Descriptor instead.
func (*DeleteTodoIn) Descriptor() ([]byte, []int) {
	return file_hyf_v1_todo_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteTodoIn) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteTodoOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTodoOut) Reset() {
	*x = DeleteTodoOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hyf_v1_todo_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTodoOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodoOut) ProtoMessage() {}

func (x *DeleteTodoOut) ProtoReflect() protoreflect.Message {
	mi := &file_hyf_v1_todo_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodoOut.ProtoReflect.Descriptor instead.
func (*DeleteTodoOut) Descriptor() ([]byte, []int) {
	return file_hyf_v1_todo_proto_rawDescGZIP(), []int{10}
}

var File_hyf_v1_todo_proto protoreflect.FileDescriptor

var file_hyf_v1_todo_proto_rawDesc = []byte{
	0x0a, 0x11, 0x68, 0x79, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72,
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
	0x4e, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x49, 0x6e, 0x22, 0x2a, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x05,
	0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x22, 0x1b, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x64, 0x6f, 0x49, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64,
	0x6f, 0x4f, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22,
	0x29, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x6e, 0x12,
	0x19, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x2a, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x4f, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x74,
	0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x29, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x64, 0x6f, 0x49, 0x6e, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64,
	0x6f, 0x22, 0x2a, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x4f,
	0x75, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x1e, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0f, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x4f, 0x75, 0x74, 0x32, 0xf1,
	0x04, 0x0a, 0x07, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x56, 0x43, 0x12, 0x68, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f,
	0x73, 0x49, 0x6e, 0x1a, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x4f, 0x75,
	0x74, 0x22, 0x41, 0x92, 0x41, 0x2d, 0x0a, 0x07, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x56, 0x43, 0x12,
	0x10, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x6f, 0x64, 0x6f,
	0x73, 0x1a, 0x10, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x6f,
	0x64, 0x6f, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x6f, 0x64, 0x6f, 0x73, 0x12, 0x6a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x12,
	0x0a, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x6e, 0x1a, 0x0b, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x64, 0x6f, 0x4f, 0x75, 0x74, 0x22, 0x46, 0x92, 0x41, 0x30, 0x0a, 0x07, 0x54,
	0x6f, 0x64, 0x6f, 0x53, 0x56, 0x43, 0x12, 0x0e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20,
	0x61, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x1a, 0x15, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20,
	0x61, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x2e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2f, 0x2a,
	0x12, 0x8a, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12,
	0x0d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x6e, 0x1a, 0x0e,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x4f, 0x75, 0x74, 0x22, 0x5d,
	0x92, 0x41, 0x46, 0x0a, 0x07, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x56, 0x43, 0x12, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x1a,
	0x28, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x20, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a,
	0x01, 0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x89, 0x01,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0d, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x6e, 0x1a, 0x0e, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x4f, 0x75, 0x74, 0x22, 0x5c, 0x92, 0x41, 0x45,
	0x0a, 0x07, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x56, 0x43, 0x12, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x1a, 0x29, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x20, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x20,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x1a, 0x09,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x77, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x6f, 0x64, 0x6f, 0x49, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x42,
	0x92, 0x41, 0x2c, 0x0a, 0x07, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x56, 0x43, 0x12, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x1a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73,
	0x2f, 0x2a, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hyf_v1_todo_proto_rawDescOnce sync.Once
	file_hyf_v1_todo_proto_rawDescData = file_hyf_v1_todo_proto_rawDesc
)

func file_hyf_v1_todo_proto_rawDescGZIP() []byte {
	file_hyf_v1_todo_proto_rawDescOnce.Do(func() {
		file_hyf_v1_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_hyf_v1_todo_proto_rawDescData)
	})
	return file_hyf_v1_todo_proto_rawDescData
}

var file_hyf_v1_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_hyf_v1_todo_proto_goTypes = []interface{}{
	(*Todo)(nil),          // 0: Todo
	(*GetTodosIn)(nil),    // 1: GetTodosIn
	(*GetTodosOut)(nil),   // 2: GetTodosOut
	(*GetTodoIn)(nil),     // 3: GetTodoIn
	(*GetTodoOut)(nil),    // 4: GetTodoOut
	(*CreateTodoIn)(nil),  // 5: CreateTodoIn
	(*CreateTodoOut)(nil), // 6: CreateTodoOut
	(*UpdateTodoIn)(nil),  // 7: UpdateTodoIn
	(*UpdateTodoOut)(nil), // 8: UpdateTodoOut
	(*DeleteTodoIn)(nil),  // 9: DeleteTodoIn
	(*DeleteTodoOut)(nil), // 10: DeleteTodoOut
	(*emptypb.Empty)(nil), // 11: google.protobuf.Empty
}
var file_hyf_v1_todo_proto_depIdxs = []int32{
	0,  // 0: GetTodosOut.todos:type_name -> Todo
	0,  // 1: GetTodoOut.todo:type_name -> Todo
	0,  // 2: CreateTodoIn.todo:type_name -> Todo
	0,  // 3: CreateTodoOut.todo:type_name -> Todo
	0,  // 4: UpdateTodoIn.todo:type_name -> Todo
	0,  // 5: UpdateTodoOut.todo:type_name -> Todo
	1,  // 6: TodoSVC.GetTodos:input_type -> GetTodosIn
	3,  // 7: TodoSVC.GetTodo:input_type -> GetTodoIn
	5,  // 8: TodoSVC.CreateTodo:input_type -> CreateTodoIn
	7,  // 9: TodoSVC.UpdateTodo:input_type -> UpdateTodoIn
	9,  // 10: TodoSVC.DeleteTodo:input_type -> DeleteTodoIn
	2,  // 11: TodoSVC.GetTodos:output_type -> GetTodosOut
	4,  // 12: TodoSVC.GetTodo:output_type -> GetTodoOut
	6,  // 13: TodoSVC.CreateTodo:output_type -> CreateTodoOut
	8,  // 14: TodoSVC.UpdateTodo:output_type -> UpdateTodoOut
	11, // 15: TodoSVC.DeleteTodo:output_type -> google.protobuf.Empty
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_hyf_v1_todo_proto_init() }
func file_hyf_v1_todo_proto_init() {
	if File_hyf_v1_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hyf_v1_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_hyf_v1_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodosIn); i {
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
		file_hyf_v1_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodosOut); i {
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
		file_hyf_v1_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoIn); i {
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
		file_hyf_v1_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoOut); i {
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
		file_hyf_v1_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoIn); i {
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
		file_hyf_v1_todo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoOut); i {
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
		file_hyf_v1_todo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTodoIn); i {
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
		file_hyf_v1_todo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTodoOut); i {
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
		file_hyf_v1_todo_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTodoIn); i {
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
		file_hyf_v1_todo_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTodoOut); i {
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
			RawDescriptor: file_hyf_v1_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hyf_v1_todo_proto_goTypes,
		DependencyIndexes: file_hyf_v1_todo_proto_depIdxs,
		MessageInfos:      file_hyf_v1_todo_proto_msgTypes,
	}.Build()
	File_hyf_v1_todo_proto = out.File
	file_hyf_v1_todo_proto_rawDesc = nil
	file_hyf_v1_todo_proto_goTypes = nil
	file_hyf_v1_todo_proto_depIdxs = nil
}
