// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: push/push.proto

package pushv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_push_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_push_push_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_push_push_proto_rawDescGZIP(), []int{0}
}

func (x *CommonResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type NowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	Text string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *NowRequest) Reset() {
	*x = NowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_push_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NowRequest) ProtoMessage() {}

func (x *NowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_push_push_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NowRequest.ProtoReflect.Descriptor instead.
func (*NowRequest) Descriptor() ([]byte, []int) {
	return file_push_push_proto_rawDescGZIP(), []int{1}
}

func (x *NowRequest) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *NowRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Date *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Id   []string               `protobuf:"bytes,3,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *ScheduleRequest) Reset() {
	*x = ScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_push_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleRequest) ProtoMessage() {}

func (x *ScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_push_push_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleRequest.ProtoReflect.Descriptor instead.
func (*ScheduleRequest) Descriptor() ([]byte, []int) {
	return file_push_push_proto_rawDescGZIP(), []int{2}
}

func (x *ScheduleRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ScheduleRequest) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *ScheduleRequest) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_push_push_proto protoreflect.FileDescriptor

var file_push_push_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x75, 0x73, 0x68, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x67, 0x73, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x30, 0x0a, 0x0a, 0x4e, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x22, 0x65, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x6a, 0x0a, 0x04, 0x50, 0x75,
	0x73, 0x68, 0x12, 0x2b, 0x0a, 0x03, 0x4e, 0x6f, 0x77, 0x12, 0x0f, 0x2e, 0x67, 0x73, 0x6d, 0x2e,
	0x4e, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x73, 0x6d,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x73,
	0x6d, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x67, 0x73, 0x6d, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x73, 0x6e, 0x69, 0x70, 0x2e, 0x70,
	0x75, 0x73, 0x68, 0x3b, 0x70, 0x75, 0x73, 0x68, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_push_push_proto_rawDescOnce sync.Once
	file_push_push_proto_rawDescData = file_push_push_proto_rawDesc
)

func file_push_push_proto_rawDescGZIP() []byte {
	file_push_push_proto_rawDescOnce.Do(func() {
		file_push_push_proto_rawDescData = protoimpl.X.CompressGZIP(file_push_push_proto_rawDescData)
	})
	return file_push_push_proto_rawDescData
}

var file_push_push_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_push_push_proto_goTypes = []any{
	(*CommonResponse)(nil),        // 0: gsm.CommonResponse
	(*NowRequest)(nil),            // 1: gsm.NowRequest
	(*ScheduleRequest)(nil),       // 2: gsm.ScheduleRequest
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_push_push_proto_depIdxs = []int32{
	3, // 0: gsm.ScheduleRequest.date:type_name -> google.protobuf.Timestamp
	1, // 1: gsm.Push.Now:input_type -> gsm.NowRequest
	2, // 2: gsm.Push.Schedule:input_type -> gsm.ScheduleRequest
	0, // 3: gsm.Push.Now:output_type -> gsm.CommonResponse
	0, // 4: gsm.Push.Schedule:output_type -> gsm.CommonResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_push_push_proto_init() }
func file_push_push_proto_init() {
	if File_push_push_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_push_push_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CommonResponse); i {
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
		file_push_push_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*NowRequest); i {
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
		file_push_push_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ScheduleRequest); i {
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
			RawDescriptor: file_push_push_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_push_push_proto_goTypes,
		DependencyIndexes: file_push_push_proto_depIdxs,
		MessageInfos:      file_push_push_proto_msgTypes,
	}.Build()
	File_push_push_proto = out.File
	file_push_push_proto_rawDesc = nil
	file_push_push_proto_goTypes = nil
	file_push_push_proto_depIdxs = nil
}
