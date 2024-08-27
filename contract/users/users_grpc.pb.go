// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: users/users.proto

package usersv1

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
	Users_Register_FullMethodName          = "/users.Users/Register"
	Users_Login_FullMethodName             = "/users.Users/Login"
	Users_HasSession_FullMethodName        = "/users.Users/HasSession"
	Users_GetAccount_FullMethodName        = "/users.Users/GetAccount"
	Users_UpdateAccountData_FullMethodName = "/users.Users/UpdateAccountData"
	Users_UpdatePassword_FullMethodName    = "/users.Users/UpdatePassword"
	Users_UpdatePhoto_FullMethodName       = "/users.Users/UpdatePhoto"
	Users_AddPhone_FullMethodName          = "/users.Users/AddPhone"
	Users_UpdatePhone_FullMethodName       = "/users.Users/UpdatePhone"
	Users_RemovePhone_FullMethodName       = "/users.Users/RemovePhone"
)

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	HasSession(ctx context.Context, in *HasSessionRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	GetAccount(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*Account, error)
	UpdateAccountData(ctx context.Context, in *UpdateAccountDataRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	UpdatePhoto(ctx context.Context, in *UpdatePhotoRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	AddPhone(ctx context.Context, in *AddPhoneRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	UpdatePhone(ctx context.Context, in *UpdatePhotoRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	RemovePhone(ctx context.Context, in *RemovePhoneRequest, opts ...grpc.CallOption) (*CommonResponse, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, Users_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, Users_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) HasSession(ctx context.Context, in *HasSessionRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Users_HasSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetAccount(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*Account, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Account)
	err := c.cc.Invoke(ctx, Users_GetAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateAccountData(ctx context.Context, in *UpdateAccountDataRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Users_UpdateAccountData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Users_UpdatePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdatePhoto(ctx context.Context, in *UpdatePhotoRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Users_UpdatePhoto_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) AddPhone(ctx context.Context, in *AddPhoneRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Users_AddPhone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdatePhone(ctx context.Context, in *UpdatePhotoRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Users_UpdatePhone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) RemovePhone(ctx context.Context, in *RemovePhoneRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Users_RemovePhone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations must embed UnimplementedUsersServer
// for forward compatibility.
type UsersServer interface {
	Register(context.Context, *RegisterRequest) (*LoginResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	HasSession(context.Context, *HasSessionRequest) (*CommonResponse, error)
	GetAccount(context.Context, *CommonRequest) (*Account, error)
	UpdateAccountData(context.Context, *UpdateAccountDataRequest) (*CommonResponse, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*CommonResponse, error)
	UpdatePhoto(context.Context, *UpdatePhotoRequest) (*CommonResponse, error)
	AddPhone(context.Context, *AddPhoneRequest) (*CommonResponse, error)
	UpdatePhone(context.Context, *UpdatePhotoRequest) (*CommonResponse, error)
	RemovePhone(context.Context, *RemovePhoneRequest) (*CommonResponse, error)
	mustEmbedUnimplementedUsersServer()
}

// UnimplementedUsersServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUsersServer struct{}

func (UnimplementedUsersServer) Register(context.Context, *RegisterRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUsersServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUsersServer) HasSession(context.Context, *HasSessionRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasSession not implemented")
}
func (UnimplementedUsersServer) GetAccount(context.Context, *CommonRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedUsersServer) UpdateAccountData(context.Context, *UpdateAccountDataRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccountData not implemented")
}
func (UnimplementedUsersServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedUsersServer) UpdatePhoto(context.Context, *UpdatePhotoRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePhoto not implemented")
}
func (UnimplementedUsersServer) AddPhone(context.Context, *AddPhoneRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPhone not implemented")
}
func (UnimplementedUsersServer) UpdatePhone(context.Context, *UpdatePhotoRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePhone not implemented")
}
func (UnimplementedUsersServer) RemovePhone(context.Context, *RemovePhoneRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePhone not implemented")
}
func (UnimplementedUsersServer) mustEmbedUnimplementedUsersServer() {}
func (UnimplementedUsersServer) testEmbeddedByValue()               {}

// UnsafeUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServer will
// result in compilation errors.
type UnsafeUsersServer interface {
	mustEmbedUnimplementedUsersServer()
}

func RegisterUsersServer(s grpc.ServiceRegistrar, srv UsersServer) {
	// If the following call pancis, it indicates UnimplementedUsersServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Users_ServiceDesc, srv)
}

func _Users_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_HasSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).HasSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_HasSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).HasSession(ctx, req.(*HasSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetAccount(ctx, req.(*CommonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdateAccountData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateAccountData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_UpdateAccountData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateAccountData(ctx, req.(*UpdateAccountDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_UpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdatePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdatePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_UpdatePhoto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdatePhoto(ctx, req.(*UpdatePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_AddPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).AddPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_AddPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).AddPhone(ctx, req.(*AddPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdatePhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdatePhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_UpdatePhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdatePhone(ctx, req.(*UpdatePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_RemovePhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).RemovePhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_RemovePhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).RemovePhone(ctx, req.(*RemovePhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Users_ServiceDesc is the grpc.ServiceDesc for Users service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Users_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Users_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Users_Login_Handler,
		},
		{
			MethodName: "HasSession",
			Handler:    _Users_HasSession_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _Users_GetAccount_Handler,
		},
		{
			MethodName: "UpdateAccountData",
			Handler:    _Users_UpdateAccountData_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _Users_UpdatePassword_Handler,
		},
		{
			MethodName: "UpdatePhoto",
			Handler:    _Users_UpdatePhoto_Handler,
		},
		{
			MethodName: "AddPhone",
			Handler:    _Users_AddPhone_Handler,
		},
		{
			MethodName: "UpdatePhone",
			Handler:    _Users_UpdatePhone_Handler,
		},
		{
			MethodName: "RemovePhone",
			Handler:    _Users_RemovePhone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users/users.proto",
}