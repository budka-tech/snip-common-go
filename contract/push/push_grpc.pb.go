// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: push/push.proto

package pushv1

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
	Push_Now_FullMethodName      = "/gsm.Push/Now"
	Push_Schedule_FullMethodName = "/gsm.Push/Schedule"
)

// PushClient is the client API for Push service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushClient interface {
	Now(ctx context.Context, in *NowRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*CommonResponse, error)
}

type pushClient struct {
	cc grpc.ClientConnInterface
}

func NewPushClient(cc grpc.ClientConnInterface) PushClient {
	return &pushClient{cc}
}

func (c *pushClient) Now(ctx context.Context, in *NowRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Push_Now_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushClient) Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, Push_Schedule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushServer is the server API for Push service.
// All implementations must embed UnimplementedPushServer
// for forward compatibility.
type PushServer interface {
	Now(context.Context, *NowRequest) (*CommonResponse, error)
	Schedule(context.Context, *ScheduleRequest) (*CommonResponse, error)
	mustEmbedUnimplementedPushServer()
}

// UnimplementedPushServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPushServer struct{}

func (UnimplementedPushServer) Now(context.Context, *NowRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Now not implemented")
}
func (UnimplementedPushServer) Schedule(context.Context, *ScheduleRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Schedule not implemented")
}
func (UnimplementedPushServer) mustEmbedUnimplementedPushServer() {}
func (UnimplementedPushServer) testEmbeddedByValue()              {}

// UnsafePushServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushServer will
// result in compilation errors.
type UnsafePushServer interface {
	mustEmbedUnimplementedPushServer()
}

func RegisterPushServer(s grpc.ServiceRegistrar, srv PushServer) {
	// If the following call pancis, it indicates UnimplementedPushServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Push_ServiceDesc, srv)
}

func _Push_Now_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).Now(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Push_Now_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).Now(ctx, req.(*NowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Push_Schedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).Schedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Push_Schedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).Schedule(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Push_ServiceDesc is the grpc.ServiceDesc for Push service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Push_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gsm.Push",
	HandlerType: (*PushServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Now",
			Handler:    _Push_Now_Handler,
		},
		{
			MethodName: "Schedule",
			Handler:    _Push_Schedule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "push/push.proto",
}
