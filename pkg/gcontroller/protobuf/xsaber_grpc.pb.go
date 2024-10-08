// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.7
// source: xsaber.proto

package gcontroller

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ServicesFunc_InsertLoginServiceInfo_FullMethodName = "/gcontroller.ServicesFunc/InsertLoginServiceInfo"
	ServicesFunc_UpdateLoginServiceInfo_FullMethodName = "/gcontroller.ServicesFunc/UpdateLoginServiceInfo"
)

// ServicesFuncClient is the client API for ServicesFunc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicesFuncClient interface {
	InsertLoginServiceInfo(ctx context.Context, in *LoginServiceInfo, opts ...grpc.CallOption) (*Response, error)
	UpdateLoginServiceInfo(ctx context.Context, opts ...grpc.CallOption) (ServicesFunc_UpdateLoginServiceInfoClient, error)
}

type servicesFuncClient struct {
	cc grpc.ClientConnInterface
}

func NewServicesFuncClient(cc grpc.ClientConnInterface) ServicesFuncClient {
	return &servicesFuncClient{cc}
}

func (c *servicesFuncClient) InsertLoginServiceInfo(ctx context.Context, in *LoginServiceInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, ServicesFunc_InsertLoginServiceInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesFuncClient) UpdateLoginServiceInfo(ctx context.Context, opts ...grpc.CallOption) (ServicesFunc_UpdateLoginServiceInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServicesFunc_ServiceDesc.Streams[0], ServicesFunc_UpdateLoginServiceInfo_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &servicesFuncUpdateLoginServiceInfoClient{stream}
	return x, nil
}

type ServicesFunc_UpdateLoginServiceInfoClient interface {
	Send(*LoginServiceInfo) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type servicesFuncUpdateLoginServiceInfoClient struct {
	grpc.ClientStream
}

func (x *servicesFuncUpdateLoginServiceInfoClient) Send(m *LoginServiceInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *servicesFuncUpdateLoginServiceInfoClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServicesFuncServer is the server API for ServicesFunc service.
// All implementations must embed UnimplementedServicesFuncServer
// for forward compatibility
type ServicesFuncServer interface {
	InsertLoginServiceInfo(context.Context, *LoginServiceInfo) (*Response, error)
	UpdateLoginServiceInfo(ServicesFunc_UpdateLoginServiceInfoServer) error
	mustEmbedUnimplementedServicesFuncServer()
}

// UnimplementedServicesFuncServer must be embedded to have forward compatible implementations.
type UnimplementedServicesFuncServer struct {
}

func (UnimplementedServicesFuncServer) InsertLoginServiceInfo(context.Context, *LoginServiceInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertLoginServiceInfo not implemented")
}
func (UnimplementedServicesFuncServer) UpdateLoginServiceInfo(ServicesFunc_UpdateLoginServiceInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateLoginServiceInfo not implemented")
}
func (UnimplementedServicesFuncServer) mustEmbedUnimplementedServicesFuncServer() {}

// UnsafeServicesFuncServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicesFuncServer will
// result in compilation errors.
type UnsafeServicesFuncServer interface {
	mustEmbedUnimplementedServicesFuncServer()
}

func RegisterServicesFuncServer(s grpc.ServiceRegistrar, srv ServicesFuncServer) {
	s.RegisterService(&ServicesFunc_ServiceDesc, srv)
}

func _ServicesFunc_InsertLoginServiceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginServiceInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesFuncServer).InsertLoginServiceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesFunc_InsertLoginServiceInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesFuncServer).InsertLoginServiceInfo(ctx, req.(*LoginServiceInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesFunc_UpdateLoginServiceInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServicesFuncServer).UpdateLoginServiceInfo(&servicesFuncUpdateLoginServiceInfoServer{stream})
}

type ServicesFunc_UpdateLoginServiceInfoServer interface {
	SendAndClose(*Response) error
	Recv() (*LoginServiceInfo, error)
	grpc.ServerStream
}

type servicesFuncUpdateLoginServiceInfoServer struct {
	grpc.ServerStream
}

func (x *servicesFuncUpdateLoginServiceInfoServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *servicesFuncUpdateLoginServiceInfoServer) Recv() (*LoginServiceInfo, error) {
	m := new(LoginServiceInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServicesFunc_ServiceDesc is the grpc.ServiceDesc for ServicesFunc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServicesFunc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gcontroller.ServicesFunc",
	HandlerType: (*ServicesFuncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertLoginServiceInfo",
			Handler:    _ServicesFunc_InsertLoginServiceInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateLoginServiceInfo",
			Handler:       _ServicesFunc_UpdateLoginServiceInfo_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "xsaber.proto",
}
