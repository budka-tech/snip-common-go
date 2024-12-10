// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: assist/assist.proto

package assistv1

import (
	context "context"
	common "github.com/budka-tech/snip-common-go/contract/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Assist_CreateChat_FullMethodName        = "/assist.Assist/CreateChat"
	Assist_ChangeChatFolder_FullMethodName  = "/assist.Assist/ChangeChatFolder"
	Assist_ChangeChatTitle_FullMethodName   = "/assist.Assist/ChangeChatTitle"
	Assist_DeleteChat_FullMethodName        = "/assist.Assist/DeleteChat"
	Assist_GetChatByID_FullMethodName       = "/assist.Assist/GetChatByID"
	Assist_GetChatsInFolder_FullMethodName  = "/assist.Assist/GetChatsInFolder"
	Assist_GetUserChats_FullMethodName      = "/assist.Assist/GetUserChats"
	Assist_CreateFolder_FullMethodName      = "/assist.Assist/CreateFolder"
	Assist_DeleteFolder_FullMethodName      = "/assist.Assist/DeleteFolder"
	Assist_GetFolderByID_FullMethodName     = "/assist.Assist/GetFolderByID"
	Assist_GetUserFolders_FullMethodName    = "/assist.Assist/GetUserFolders"
	Assist_SendMessage_FullMethodName       = "/assist.Assist/SendMessage"
	Assist_GetMessageByID_FullMethodName    = "/assist.Assist/GetMessageByID"
	Assist_GetMessagesInChat_FullMethodName = "/assist.Assist/GetMessagesInChat"
)

// AssistClient is the client API for Assist service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssistClient interface {
	CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error)
	ChangeChatFolder(ctx context.Context, in *ChangeChatFolderRequest, opts ...grpc.CallOption) (*ChangeChatFolderResponse, error)
	ChangeChatTitle(ctx context.Context, in *ChangeChatTitleRequest, opts ...grpc.CallOption) (*ChangeChatTitleResponse, error)
	DeleteChat(ctx context.Context, in *DeleteChatRequest, opts ...grpc.CallOption) (*common.Response, error)
	GetChatByID(ctx context.Context, in *GetChatByIDRequest, opts ...grpc.CallOption) (*GetChatByIDResponse, error)
	GetChatsInFolder(ctx context.Context, in *GetChatsInFolderRequest, opts ...grpc.CallOption) (*GetChatsInFolderResponse, error)
	GetUserChats(ctx context.Context, in *GetUserChatsRequest, opts ...grpc.CallOption) (*GetUserChatsResponse, error)
	CreateFolder(ctx context.Context, in *CreateFolderRequest, opts ...grpc.CallOption) (*CreateFolderResponse, error)
	DeleteFolder(ctx context.Context, in *DeleteFolderRequest, opts ...grpc.CallOption) (*common.Response, error)
	GetFolderByID(ctx context.Context, in *GetFolderByIDRequest, opts ...grpc.CallOption) (*GetFolderByIDResponse, error)
	GetUserFolders(ctx context.Context, in *GetUserFoldersRequest, opts ...grpc.CallOption) (*GetUserFoldersResponse, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	GetMessageByID(ctx context.Context, in *GetMessageByIDRequest, opts ...grpc.CallOption) (*GetMessageByIDResponse, error)
	GetMessagesInChat(ctx context.Context, in *GetMessagesInChatRequest, opts ...grpc.CallOption) (*GetMessagesInChatResponse, error)
}

type assistClient struct {
	cc grpc.ClientConnInterface
}

func NewAssistClient(cc grpc.ClientConnInterface) AssistClient {
	return &assistClient{cc}
}

func (c *assistClient) CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateChatResponse)
	err := c.cc.Invoke(ctx, Assist_CreateChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistClient) ChangeChatFolder(ctx context.Context, in *ChangeChatFolderRequest, opts ...grpc.CallOption) (*ChangeChatFolderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangeChatFolderResponse)
	err := c.cc.Invoke(ctx, Assist_ChangeChatFolder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistClient) ChangeChatTitle(ctx context.Context, in *ChangeChatTitleRequest, opts ...grpc.CallOption) (*ChangeChatTitleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangeChatTitleResponse)
	err := c.cc.Invoke(ctx, Assist_ChangeChatTitle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistClient) DeleteChat(ctx context.Context, in *DeleteChatRequest, opts ...grpc.CallOption) (*common.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Response)
	err := c.cc.Invoke(ctx, Assist_DeleteChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistClient) GetChatByID(ctx context.Context, in *GetChatByIDRequest, opts ...grpc.CallOption) (*GetChatByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChatByIDResponse)
	err := c.cc.Invoke(ctx, Assist_GetChatByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistClient) GetChatsInFolder(ctx context.Context, in *GetChatsInFolderRequest, opts ...grpc.CallOption) (*GetChatsInFolderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChatsInFolderResponse)
	err := c.cc.Invoke(ctx, Assist_GetChatsInFolder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistClient) GetUserChats(ctx context.Context, in *GetUserChatsRequest, opts ...grpc.CallOption) (*GetUserChatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserChatsResponse)
	err := c.cc.Invoke(ctx, Assist_GetUserChats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistClient) CreateFolder(ctx context.Context, in *CreateFolderRequest, opts ...grpc.CallOption) (*CreateFolderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateFolderResponse)
	err := c.cc.Invoke(ctx, Assist_CreateFolder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistClient) DeleteFolder(ctx context.Context, in *DeleteFolderRequest, opts ...grpc.CallOption) (*common.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Response)
	err := c.cc.Invoke(ctx, Assist_DeleteFolder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistClient) GetFolderByID(ctx context.Context, in *GetFolderByIDRequest, opts ...grpc.CallOption) (*GetFolderByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFolderByIDResponse)
	err := c.cc.Invoke(ctx, Assist_GetFolderByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistClient) GetUserFolders(ctx context.Context, in *GetUserFoldersRequest, opts ...grpc.CallOption) (*GetUserFoldersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserFoldersResponse)
	err := c.cc.Invoke(ctx, Assist_GetUserFolders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, Assist_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistClient) GetMessageByID(ctx context.Context, in *GetMessageByIDRequest, opts ...grpc.CallOption) (*GetMessageByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMessageByIDResponse)
	err := c.cc.Invoke(ctx, Assist_GetMessageByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistClient) GetMessagesInChat(ctx context.Context, in *GetMessagesInChatRequest, opts ...grpc.CallOption) (*GetMessagesInChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMessagesInChatResponse)
	err := c.cc.Invoke(ctx, Assist_GetMessagesInChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssistServer is the server API for Assist service.
// All implementations must embed UnimplementedAssistServer
// for forward compatibility.
type AssistServer interface {
	CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error)
	ChangeChatFolder(context.Context, *ChangeChatFolderRequest) (*ChangeChatFolderResponse, error)
	ChangeChatTitle(context.Context, *ChangeChatTitleRequest) (*ChangeChatTitleResponse, error)
	DeleteChat(context.Context, *DeleteChatRequest) (*common.Response, error)
	GetChatByID(context.Context, *GetChatByIDRequest) (*GetChatByIDResponse, error)
	GetChatsInFolder(context.Context, *GetChatsInFolderRequest) (*GetChatsInFolderResponse, error)
	GetUserChats(context.Context, *GetUserChatsRequest) (*GetUserChatsResponse, error)
	CreateFolder(context.Context, *CreateFolderRequest) (*CreateFolderResponse, error)
	DeleteFolder(context.Context, *DeleteFolderRequest) (*common.Response, error)
	GetFolderByID(context.Context, *GetFolderByIDRequest) (*GetFolderByIDResponse, error)
	GetUserFolders(context.Context, *GetUserFoldersRequest) (*GetUserFoldersResponse, error)
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	GetMessageByID(context.Context, *GetMessageByIDRequest) (*GetMessageByIDResponse, error)
	GetMessagesInChat(context.Context, *GetMessagesInChatRequest) (*GetMessagesInChatResponse, error)
	mustEmbedUnimplementedAssistServer()
}

// UnimplementedAssistServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAssistServer struct{}

func (UnimplementedAssistServer) CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedAssistServer) ChangeChatFolder(context.Context, *ChangeChatFolderRequest) (*ChangeChatFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeChatFolder not implemented")
}
func (UnimplementedAssistServer) ChangeChatTitle(context.Context, *ChangeChatTitleRequest) (*ChangeChatTitleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeChatTitle not implemented")
}
func (UnimplementedAssistServer) DeleteChat(context.Context, *DeleteChatRequest) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChat not implemented")
}
func (UnimplementedAssistServer) GetChatByID(context.Context, *GetChatByIDRequest) (*GetChatByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatByID not implemented")
}
func (UnimplementedAssistServer) GetChatsInFolder(context.Context, *GetChatsInFolderRequest) (*GetChatsInFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatsInFolder not implemented")
}
func (UnimplementedAssistServer) GetUserChats(context.Context, *GetUserChatsRequest) (*GetUserChatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserChats not implemented")
}
func (UnimplementedAssistServer) CreateFolder(context.Context, *CreateFolderRequest) (*CreateFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFolder not implemented")
}
func (UnimplementedAssistServer) DeleteFolder(context.Context, *DeleteFolderRequest) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFolder not implemented")
}
func (UnimplementedAssistServer) GetFolderByID(context.Context, *GetFolderByIDRequest) (*GetFolderByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFolderByID not implemented")
}
func (UnimplementedAssistServer) GetUserFolders(context.Context, *GetUserFoldersRequest) (*GetUserFoldersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFolders not implemented")
}
func (UnimplementedAssistServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedAssistServer) GetMessageByID(context.Context, *GetMessageByIDRequest) (*GetMessageByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageByID not implemented")
}
func (UnimplementedAssistServer) GetMessagesInChat(context.Context, *GetMessagesInChatRequest) (*GetMessagesInChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessagesInChat not implemented")
}
func (UnimplementedAssistServer) mustEmbedUnimplementedAssistServer() {}
func (UnimplementedAssistServer) testEmbeddedByValue()                {}

// UnsafeAssistServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssistServer will
// result in compilation errors.
type UnsafeAssistServer interface {
	mustEmbedUnimplementedAssistServer()
}

func RegisterAssistServer(s grpc.ServiceRegistrar, srv AssistServer) {
	// If the following call pancis, it indicates UnimplementedAssistServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Assist_ServiceDesc, srv)
}

func _Assist_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assist_CreateChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistServer).CreateChat(ctx, req.(*CreateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assist_ChangeChatFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeChatFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistServer).ChangeChatFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assist_ChangeChatFolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistServer).ChangeChatFolder(ctx, req.(*ChangeChatFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assist_ChangeChatTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeChatTitleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistServer).ChangeChatTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assist_ChangeChatTitle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistServer).ChangeChatTitle(ctx, req.(*ChangeChatTitleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assist_DeleteChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistServer).DeleteChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assist_DeleteChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistServer).DeleteChat(ctx, req.(*DeleteChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assist_GetChatByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistServer).GetChatByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assist_GetChatByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistServer).GetChatByID(ctx, req.(*GetChatByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assist_GetChatsInFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatsInFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistServer).GetChatsInFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assist_GetChatsInFolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistServer).GetChatsInFolder(ctx, req.(*GetChatsInFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assist_GetUserChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserChatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistServer).GetUserChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assist_GetUserChats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistServer).GetUserChats(ctx, req.(*GetUserChatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assist_CreateFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistServer).CreateFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assist_CreateFolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistServer).CreateFolder(ctx, req.(*CreateFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assist_DeleteFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistServer).DeleteFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assist_DeleteFolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistServer).DeleteFolder(ctx, req.(*DeleteFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assist_GetFolderByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFolderByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistServer).GetFolderByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assist_GetFolderByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistServer).GetFolderByID(ctx, req.(*GetFolderByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assist_GetUserFolders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFoldersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistServer).GetUserFolders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assist_GetUserFolders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistServer).GetUserFolders(ctx, req.(*GetUserFoldersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assist_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assist_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assist_GetMessageByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistServer).GetMessageByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assist_GetMessageByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistServer).GetMessageByID(ctx, req.(*GetMessageByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Assist_GetMessagesInChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesInChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistServer).GetMessagesInChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Assist_GetMessagesInChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistServer).GetMessagesInChat(ctx, req.(*GetMessagesInChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Assist_ServiceDesc is the grpc.ServiceDesc for Assist service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Assist_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "assist.Assist",
	HandlerType: (*AssistServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChat",
			Handler:    _Assist_CreateChat_Handler,
		},
		{
			MethodName: "ChangeChatFolder",
			Handler:    _Assist_ChangeChatFolder_Handler,
		},
		{
			MethodName: "ChangeChatTitle",
			Handler:    _Assist_ChangeChatTitle_Handler,
		},
		{
			MethodName: "DeleteChat",
			Handler:    _Assist_DeleteChat_Handler,
		},
		{
			MethodName: "GetChatByID",
			Handler:    _Assist_GetChatByID_Handler,
		},
		{
			MethodName: "GetChatsInFolder",
			Handler:    _Assist_GetChatsInFolder_Handler,
		},
		{
			MethodName: "GetUserChats",
			Handler:    _Assist_GetUserChats_Handler,
		},
		{
			MethodName: "CreateFolder",
			Handler:    _Assist_CreateFolder_Handler,
		},
		{
			MethodName: "DeleteFolder",
			Handler:    _Assist_DeleteFolder_Handler,
		},
		{
			MethodName: "GetFolderByID",
			Handler:    _Assist_GetFolderByID_Handler,
		},
		{
			MethodName: "GetUserFolders",
			Handler:    _Assist_GetUserFolders_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Assist_SendMessage_Handler,
		},
		{
			MethodName: "GetMessageByID",
			Handler:    _Assist_GetMessageByID_Handler,
		},
		{
			MethodName: "GetMessagesInChat",
			Handler:    _Assist_GetMessagesInChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "assist/assist.proto",
}
