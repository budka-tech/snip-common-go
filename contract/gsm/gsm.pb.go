// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.3
// source: gsm/gsm.proto

package gsmv1

import (
	common "github.com/budka-tech/snip-common-go/contract/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Method int32

const (
	Method_FlASH_CALL      Method = 0
	Method_FlASH_CALL_BACK Method = 1
)

// Enum value maps for Method.
var (
	Method_name = map[int32]string{
		0: "FlASH_CALL",
		1: "FlASH_CALL_BACK",
	}
	Method_value = map[string]int32{
		"FlASH_CALL":      0,
		"FlASH_CALL_BACK": 1,
	}
)

func (x Method) Enum() *Method {
	p := new(Method)
	*p = x
	return p
}

func (x Method) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Method) Descriptor() protoreflect.EnumDescriptor {
	return file_gsm_gsm_proto_enumTypes[0].Descriptor()
}

func (Method) Type() protoreflect.EnumType {
	return &file_gsm_gsm_proto_enumTypes[0]
}

func (x Method) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Method.Descriptor instead.
func (Method) EnumDescriptor() ([]byte, []int) {
	return file_gsm_gsm_proto_rawDescGZIP(), []int{0}
}

type FlashCallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Code  string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *FlashCallRequest) Reset() {
	*x = FlashCallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsm_gsm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlashCallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlashCallRequest) ProtoMessage() {}

func (x *FlashCallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gsm_gsm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlashCallRequest.ProtoReflect.Descriptor instead.
func (*FlashCallRequest) Descriptor() ([]byte, []int) {
	return file_gsm_gsm_proto_rawDescGZIP(), []int{0}
}

func (x *FlashCallRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *FlashCallRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type FlashCallBackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *FlashCallBackRequest) Reset() {
	*x = FlashCallBackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gsm_gsm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlashCallBackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlashCallBackRequest) ProtoMessage() {}

func (x *FlashCallBackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gsm_gsm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlashCallBackRequest.ProtoReflect.Descriptor instead.
func (*FlashCallBackRequest) Descriptor() ([]byte, []int) {
	return file_gsm_gsm_proto_rawDescGZIP(), []int{1}
}

func (x *FlashCallBackRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_gsm_gsm_proto protoreflect.FileDescriptor

var file_gsm_gsm_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x73, 0x6d, 0x2f, 0x67, 0x73, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x67, 0x73, 0x6d, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x10, 0x46, 0x6c, 0x61,
	0x73, 0x68, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x2c, 0x0a, 0x14, 0x46, 0x6c, 0x61, 0x73, 0x68,
	0x43, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2a, 0x2d, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x0e, 0x0a, 0x0a, 0x46, 0x6c, 0x41, 0x53, 0x48, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x46, 0x6c, 0x41, 0x53, 0x48, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x42, 0x41,
	0x43, 0x4b, 0x10, 0x01, 0x32, 0x79, 0x0a, 0x03, 0x47, 0x73, 0x6d, 0x12, 0x34, 0x0a, 0x09, 0x46,
	0x6c, 0x61, 0x73, 0x68, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x67, 0x73, 0x6d, 0x2e, 0x46,
	0x6c, 0x61, 0x73, 0x68, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x43, 0x61, 0x6c, 0x6c, 0x42, 0x61,
	0x63, 0x6b, 0x12, 0x19, 0x2e, 0x67, 0x73, 0x6d, 0x2e, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x43, 0x61,
	0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x10, 0x5a, 0x0e, 0x73, 0x6e, 0x69, 0x70, 0x2e, 0x67, 0x73, 0x6d, 0x3b, 0x67, 0x73, 0x6d, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gsm_gsm_proto_rawDescOnce sync.Once
	file_gsm_gsm_proto_rawDescData = file_gsm_gsm_proto_rawDesc
)

func file_gsm_gsm_proto_rawDescGZIP() []byte {
	file_gsm_gsm_proto_rawDescOnce.Do(func() {
		file_gsm_gsm_proto_rawDescData = protoimpl.X.CompressGZIP(file_gsm_gsm_proto_rawDescData)
	})
	return file_gsm_gsm_proto_rawDescData
}

var file_gsm_gsm_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gsm_gsm_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gsm_gsm_proto_goTypes = []any{
	(Method)(0),                  // 0: gsm.Method
	(*FlashCallRequest)(nil),     // 1: gsm.FlashCallRequest
	(*FlashCallBackRequest)(nil), // 2: gsm.FlashCallBackRequest
	(*common.Response)(nil),      // 3: common.Response
}
var file_gsm_gsm_proto_depIdxs = []int32{
	1, // 0: gsm.Gsm.FlashCall:input_type -> gsm.FlashCallRequest
	2, // 1: gsm.Gsm.FlashCallBack:input_type -> gsm.FlashCallBackRequest
	3, // 2: gsm.Gsm.FlashCall:output_type -> common.Response
	3, // 3: gsm.Gsm.FlashCallBack:output_type -> common.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gsm_gsm_proto_init() }
func file_gsm_gsm_proto_init() {
	if File_gsm_gsm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gsm_gsm_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FlashCallRequest); i {
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
		file_gsm_gsm_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FlashCallBackRequest); i {
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
			RawDescriptor: file_gsm_gsm_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gsm_gsm_proto_goTypes,
		DependencyIndexes: file_gsm_gsm_proto_depIdxs,
		EnumInfos:         file_gsm_gsm_proto_enumTypes,
		MessageInfos:      file_gsm_gsm_proto_msgTypes,
	}.Build()
	File_gsm_gsm_proto = out.File
	file_gsm_gsm_proto_rawDesc = nil
	file_gsm_gsm_proto_goTypes = nil
	file_gsm_gsm_proto_depIdxs = nil
}
