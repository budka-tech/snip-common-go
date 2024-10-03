// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: server/server.proto

package serverv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageType int32

const (
	MessageType_Call        MessageType = 0
	MessageType_Subscribe   MessageType = 1
	MessageType_Unsubscribe MessageType = 2
	MessageType_Event       MessageType = 3
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "Call",
		1: "Subscribe",
		2: "Unsubscribe",
		3: "Event",
	}
	MessageType_value = map[string]int32{
		"Call":        0,
		"Subscribe":   1,
		"Unsubscribe": 2,
		"Event":       3,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_server_server_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_server_server_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{0}
}

type Domain int32

const (
	Domain_Auth   Domain = 0
	Domain_Users  Domain = 1
	Domain_Assist Domain = 2
)

// Enum value maps for Domain.
var (
	Domain_name = map[int32]string{
		0: "Auth",
		1: "Users",
		2: "Assist",
	}
	Domain_value = map[string]int32{
		"Auth":   0,
		"Users":  1,
		"Assist": 2,
	}
)

func (x Domain) Enum() *Domain {
	p := new(Domain)
	*p = x
	return p
}

func (x Domain) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Domain) Descriptor() protoreflect.EnumDescriptor {
	return file_server_server_proto_enumTypes[1].Descriptor()
}

func (Domain) Type() protoreflect.EnumType {
	return &file_server_server_proto_enumTypes[1]
}

func (x Domain) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Domain.Descriptor instead.
func (Domain) EnumDescriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{1}
}

type SubscribeAction int32

const (
	SubscribeAction_Init   SubscribeAction = 0
	SubscribeAction_Create SubscribeAction = 1
	SubscribeAction_Update SubscribeAction = 2
	SubscribeAction_Delete SubscribeAction = 3
)

// Enum value maps for SubscribeAction.
var (
	SubscribeAction_name = map[int32]string{
		0: "Init",
		1: "Create",
		2: "Update",
		3: "Delete",
	}
	SubscribeAction_value = map[string]int32{
		"Init":   0,
		"Create": 1,
		"Update": 2,
		"Delete": 3,
	}
)

func (x SubscribeAction) Enum() *SubscribeAction {
	p := new(SubscribeAction)
	*p = x
	return p
}

func (x SubscribeAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubscribeAction) Descriptor() protoreflect.EnumDescriptor {
	return file_server_server_proto_enumTypes[2].Descriptor()
}

func (SubscribeAction) Type() protoreflect.EnumType {
	return &file_server_server_proto_enumTypes[2]
}

func (x SubscribeAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubscribeAction.Descriptor instead.
func (SubscribeAction) EnumDescriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{2}
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=server.MessageType" json:"type,omitempty"`
	Domain Domain      `protobuf:"varint,2,opt,name=domain,proto3,enum=server.Domain" json:"domain,omitempty"`
	Id     uint32      `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_Call
}

func (x *Meta) GetDomain() Domain {
	if x != nil {
		return x.Domain
	}
	return Domain_Auth
}

func (x *Meta) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ParamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Meta      `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Data *anypb.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ParamsRequest) Reset() {
	*x = ParamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamsRequest) ProtoMessage() {}

func (x *ParamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamsRequest.ProtoReflect.Descriptor instead.
func (*ParamsRequest) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{1}
}

func (x *ParamsRequest) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ParamsRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta   *Meta      `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Status uint32     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Data   *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Response) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Response) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type SubscribeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action SubscribeAction `protobuf:"varint,1,opt,name=action,proto3,enum=server.SubscribeAction" json:"action,omitempty"`
	Data   *anypb.Any      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SubscribeData) Reset() {
	*x = SubscribeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeData) ProtoMessage() {}

func (x *SubscribeData) ProtoReflect() protoreflect.Message {
	mi := &file_server_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeData.ProtoReflect.Descriptor instead.
func (*SubscribeData) Descriptor() ([]byte, []int) {
	return file_server_server_proto_rawDescGZIP(), []int{3}
}

func (x *SubscribeData) GetAction() SubscribeAction {
	if x != nil {
		return x.Action
	}
	return SubscribeAction_Init
}

func (x *SubscribeData) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_server_server_proto protoreflect.FileDescriptor

var file_server_server_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x27, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5b,
	0x0a, 0x0d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6e, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6a, 0x0a, 0x0d, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x42, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x03, 0x2a, 0x29, 0x0a, 0x06, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x10, 0x02, 0x2a, 0x3f, 0x0a, 0x0f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x6e, 0x69,
	0x74, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x03, 0x32, 0x42, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x38, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x16, 0x5a, 0x14, 0x73,
	0x6e, 0x69, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_server_proto_rawDescOnce sync.Once
	file_server_server_proto_rawDescData = file_server_server_proto_rawDesc
)

func file_server_server_proto_rawDescGZIP() []byte {
	file_server_server_proto_rawDescOnce.Do(func() {
		file_server_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_server_proto_rawDescData)
	})
	return file_server_server_proto_rawDescData
}

var file_server_server_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_server_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_server_server_proto_goTypes = []any{
	(MessageType)(0),      // 0: server.MessageType
	(Domain)(0),           // 1: server.Domain
	(SubscribeAction)(0),  // 2: server.SubscribeAction
	(*Meta)(nil),          // 3: server.Meta
	(*ParamsRequest)(nil), // 4: server.ParamsRequest
	(*Response)(nil),      // 5: server.Response
	(*SubscribeData)(nil), // 6: server.SubscribeData
	(*anypb.Any)(nil),     // 7: google.protobuf.Any
	(*emptypb.Empty)(nil), // 8: google.protobuf.Empty
}
var file_server_server_proto_depIdxs = []int32{
	0, // 0: server.Meta.type:type_name -> server.MessageType
	1, // 1: server.Meta.domain:type_name -> server.Domain
	3, // 2: server.ParamsRequest.meta:type_name -> server.Meta
	7, // 3: server.ParamsRequest.data:type_name -> google.protobuf.Any
	3, // 4: server.Response.meta:type_name -> server.Meta
	7, // 5: server.Response.data:type_name -> google.protobuf.Any
	2, // 6: server.SubscribeData.action:type_name -> server.SubscribeAction
	7, // 7: server.SubscribeData.data:type_name -> google.protobuf.Any
	4, // 8: server.Server.Request:input_type -> server.ParamsRequest
	8, // 9: server.Server.Request:output_type -> google.protobuf.Empty
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_server_server_proto_init() }
func file_server_server_proto_init() {
	if File_server_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_server_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Meta); i {
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
		file_server_server_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ParamsRequest); i {
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
		file_server_server_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Response); i {
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
		file_server_server_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SubscribeData); i {
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
			RawDescriptor: file_server_server_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_server_proto_goTypes,
		DependencyIndexes: file_server_server_proto_depIdxs,
		EnumInfos:         file_server_server_proto_enumTypes,
		MessageInfos:      file_server_server_proto_msgTypes,
	}.Build()
	File_server_server_proto = out.File
	file_server_server_proto_rawDesc = nil
	file_server_server_proto_goTypes = nil
	file_server_server_proto_depIdxs = nil
}
