// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: users/users.proto

package usersv1

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
	Users_Identify_FullMethodName          = "/users.Users/Identify"
	Users_Login_FullMethodName             = "/users.Users/Login"
	Users_Logout_FullMethodName            = "/users.Users/Logout"
	Users_CheckCode_FullMethodName         = "/users.Users/CheckCode"
	Users_GetAccount_FullMethodName        = "/users.Users/GetAccount"
	Users_GetAccountShort_FullMethodName   = "/users.Users/GetAccountShort"
	Users_GetSessions_FullMethodName       = "/users.Users/GetSessions"
	Users_UpdateAccountData_FullMethodName = "/users.Users/UpdateAccountData"
	Users_GetPhones_FullMethodName         = "/users.Users/GetPhones"
	Users_AttachPhone_FullMethodName       = "/users.Users/AttachPhone"
	Users_DetachPhone_FullMethodName       = "/users.Users/DetachPhone"
	Users_SetPrimaryPhone_FullMethodName   = "/users.Users/SetPrimaryPhone"
	Users_GetEmails_FullMethodName         = "/users.Users/GetEmails"
	Users_AttachEmail_FullMethodName       = "/users.Users/AttachEmail"
	Users_DetachEmail_FullMethodName       = "/users.Users/DetachEmail"
	Users_SetPrimaryEmail_FullMethodName   = "/users.Users/SetPrimaryEmail"
)

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	Identify(ctx context.Context, in *IdentifyRequest, opts ...grpc.CallOption) (*IdentifyResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*common.Response, error)
	CheckCode(ctx context.Context, in *CheckCodeRequest, opts ...grpc.CallOption) (*common.Response, error)
	GetAccount(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*Account, error)
	GetAccountShort(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*Account, error)
	GetSessions(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*GetSessionsResponse, error)
	UpdateAccountData(ctx context.Context, in *UpdateAccountDataRequest, opts ...grpc.CallOption) (*common.Response, error)
	GetPhones(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*GetPhonesResponse, error)
	AttachPhone(ctx context.Context, in *PhoneManipulationRequest, opts ...grpc.CallOption) (*ConfirmResponse, error)
	DetachPhone(ctx context.Context, in *PhoneManipulationRequest, opts ...grpc.CallOption) (*ConfirmResponse, error)
	SetPrimaryPhone(ctx context.Context, in *PhoneManipulationRequest, opts ...grpc.CallOption) (*ConfirmResponse, error)
	GetEmails(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*GetEmailsResponse, error)
	AttachEmail(ctx context.Context, in *EmailManipulationRequest, opts ...grpc.CallOption) (*ConfirmResponse, error)
	DetachEmail(ctx context.Context, in *EmailManipulationRequest, opts ...grpc.CallOption) (*ConfirmResponse, error)
	SetPrimaryEmail(ctx context.Context, in *EmailManipulationRequest, opts ...grpc.CallOption) (*ConfirmResponse, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Identify(ctx context.Context, in *IdentifyRequest, opts ...grpc.CallOption) (*IdentifyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IdentifyResponse)
	err := c.cc.Invoke(ctx, Users_Identify_FullMethodName, in, out, cOpts...)
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

func (c *usersClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*common.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Response)
	err := c.cc.Invoke(ctx, Users_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) CheckCode(ctx context.Context, in *CheckCodeRequest, opts ...grpc.CallOption) (*common.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Response)
	err := c.cc.Invoke(ctx, Users_CheckCode_FullMethodName, in, out, cOpts...)
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

func (c *usersClient) GetAccountShort(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*Account, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Account)
	err := c.cc.Invoke(ctx, Users_GetAccountShort_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetSessions(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*GetSessionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSessionsResponse)
	err := c.cc.Invoke(ctx, Users_GetSessions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateAccountData(ctx context.Context, in *UpdateAccountDataRequest, opts ...grpc.CallOption) (*common.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.Response)
	err := c.cc.Invoke(ctx, Users_UpdateAccountData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetPhones(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*GetPhonesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPhonesResponse)
	err := c.cc.Invoke(ctx, Users_GetPhones_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) AttachPhone(ctx context.Context, in *PhoneManipulationRequest, opts ...grpc.CallOption) (*ConfirmResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfirmResponse)
	err := c.cc.Invoke(ctx, Users_AttachPhone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) DetachPhone(ctx context.Context, in *PhoneManipulationRequest, opts ...grpc.CallOption) (*ConfirmResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfirmResponse)
	err := c.cc.Invoke(ctx, Users_DetachPhone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) SetPrimaryPhone(ctx context.Context, in *PhoneManipulationRequest, opts ...grpc.CallOption) (*ConfirmResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfirmResponse)
	err := c.cc.Invoke(ctx, Users_SetPrimaryPhone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetEmails(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*GetEmailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEmailsResponse)
	err := c.cc.Invoke(ctx, Users_GetEmails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) AttachEmail(ctx context.Context, in *EmailManipulationRequest, opts ...grpc.CallOption) (*ConfirmResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfirmResponse)
	err := c.cc.Invoke(ctx, Users_AttachEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) DetachEmail(ctx context.Context, in *EmailManipulationRequest, opts ...grpc.CallOption) (*ConfirmResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfirmResponse)
	err := c.cc.Invoke(ctx, Users_DetachEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) SetPrimaryEmail(ctx context.Context, in *EmailManipulationRequest, opts ...grpc.CallOption) (*ConfirmResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfirmResponse)
	err := c.cc.Invoke(ctx, Users_SetPrimaryEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations must embed UnimplementedUsersServer
// for forward compatibility.
type UsersServer interface {
	Identify(context.Context, *IdentifyRequest) (*IdentifyResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*common.Response, error)
	CheckCode(context.Context, *CheckCodeRequest) (*common.Response, error)
	GetAccount(context.Context, *CommonRequest) (*Account, error)
	GetAccountShort(context.Context, *CommonRequest) (*Account, error)
	GetSessions(context.Context, *CommonRequest) (*GetSessionsResponse, error)
	UpdateAccountData(context.Context, *UpdateAccountDataRequest) (*common.Response, error)
	GetPhones(context.Context, *CommonRequest) (*GetPhonesResponse, error)
	AttachPhone(context.Context, *PhoneManipulationRequest) (*ConfirmResponse, error)
	DetachPhone(context.Context, *PhoneManipulationRequest) (*ConfirmResponse, error)
	SetPrimaryPhone(context.Context, *PhoneManipulationRequest) (*ConfirmResponse, error)
	GetEmails(context.Context, *CommonRequest) (*GetEmailsResponse, error)
	AttachEmail(context.Context, *EmailManipulationRequest) (*ConfirmResponse, error)
	DetachEmail(context.Context, *EmailManipulationRequest) (*ConfirmResponse, error)
	SetPrimaryEmail(context.Context, *EmailManipulationRequest) (*ConfirmResponse, error)
	mustEmbedUnimplementedUsersServer()
}

// UnimplementedUsersServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUsersServer struct{}

func (UnimplementedUsersServer) Identify(context.Context, *IdentifyRequest) (*IdentifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Identify not implemented")
}
func (UnimplementedUsersServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUsersServer) Logout(context.Context, *LogoutRequest) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUsersServer) CheckCode(context.Context, *CheckCodeRequest) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCode not implemented")
}
func (UnimplementedUsersServer) GetAccount(context.Context, *CommonRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedUsersServer) GetAccountShort(context.Context, *CommonRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountShort not implemented")
}
func (UnimplementedUsersServer) GetSessions(context.Context, *CommonRequest) (*GetSessionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessions not implemented")
}
func (UnimplementedUsersServer) UpdateAccountData(context.Context, *UpdateAccountDataRequest) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccountData not implemented")
}
func (UnimplementedUsersServer) GetPhones(context.Context, *CommonRequest) (*GetPhonesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPhones not implemented")
}
func (UnimplementedUsersServer) AttachPhone(context.Context, *PhoneManipulationRequest) (*ConfirmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachPhone not implemented")
}
func (UnimplementedUsersServer) DetachPhone(context.Context, *PhoneManipulationRequest) (*ConfirmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetachPhone not implemented")
}
func (UnimplementedUsersServer) SetPrimaryPhone(context.Context, *PhoneManipulationRequest) (*ConfirmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPrimaryPhone not implemented")
}
func (UnimplementedUsersServer) GetEmails(context.Context, *CommonRequest) (*GetEmailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmails not implemented")
}
func (UnimplementedUsersServer) AttachEmail(context.Context, *EmailManipulationRequest) (*ConfirmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachEmail not implemented")
}
func (UnimplementedUsersServer) DetachEmail(context.Context, *EmailManipulationRequest) (*ConfirmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetachEmail not implemented")
}
func (UnimplementedUsersServer) SetPrimaryEmail(context.Context, *EmailManipulationRequest) (*ConfirmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPrimaryEmail not implemented")
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

func _Users_Identify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Identify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_Identify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Identify(ctx, req.(*IdentifyRequest))
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

func _Users_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_CheckCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CheckCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_CheckCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CheckCode(ctx, req.(*CheckCodeRequest))
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

func _Users_GetAccountShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetAccountShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetAccountShort_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetAccountShort(ctx, req.(*CommonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetSessions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetSessions(ctx, req.(*CommonRequest))
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

func _Users_GetPhones_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetPhones(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetPhones_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetPhones(ctx, req.(*CommonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_AttachPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneManipulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).AttachPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_AttachPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).AttachPhone(ctx, req.(*PhoneManipulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_DetachPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneManipulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).DetachPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_DetachPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).DetachPhone(ctx, req.(*PhoneManipulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_SetPrimaryPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneManipulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).SetPrimaryPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_SetPrimaryPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).SetPrimaryPhone(ctx, req.(*PhoneManipulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetEmails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetEmails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetEmails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetEmails(ctx, req.(*CommonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_AttachEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailManipulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).AttachEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_AttachEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).AttachEmail(ctx, req.(*EmailManipulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_DetachEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailManipulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).DetachEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_DetachEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).DetachEmail(ctx, req.(*EmailManipulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_SetPrimaryEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailManipulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).SetPrimaryEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_SetPrimaryEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).SetPrimaryEmail(ctx, req.(*EmailManipulationRequest))
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
			MethodName: "Identify",
			Handler:    _Users_Identify_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Users_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Users_Logout_Handler,
		},
		{
			MethodName: "CheckCode",
			Handler:    _Users_CheckCode_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _Users_GetAccount_Handler,
		},
		{
			MethodName: "GetAccountShort",
			Handler:    _Users_GetAccountShort_Handler,
		},
		{
			MethodName: "GetSessions",
			Handler:    _Users_GetSessions_Handler,
		},
		{
			MethodName: "UpdateAccountData",
			Handler:    _Users_UpdateAccountData_Handler,
		},
		{
			MethodName: "GetPhones",
			Handler:    _Users_GetPhones_Handler,
		},
		{
			MethodName: "AttachPhone",
			Handler:    _Users_AttachPhone_Handler,
		},
		{
			MethodName: "DetachPhone",
			Handler:    _Users_DetachPhone_Handler,
		},
		{
			MethodName: "SetPrimaryPhone",
			Handler:    _Users_SetPrimaryPhone_Handler,
		},
		{
			MethodName: "GetEmails",
			Handler:    _Users_GetEmails_Handler,
		},
		{
			MethodName: "AttachEmail",
			Handler:    _Users_AttachEmail_Handler,
		},
		{
			MethodName: "DetachEmail",
			Handler:    _Users_DetachEmail_Handler,
		},
		{
			MethodName: "SetPrimaryEmail",
			Handler:    _Users_SetPrimaryEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users/users.proto",
}
