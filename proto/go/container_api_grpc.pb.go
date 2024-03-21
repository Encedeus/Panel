// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: container_api.proto

package protoapi

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

// ContainerClient is the client API for Container service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContainerClient interface {
	Create(ctx context.Context, in *ContainerCreateParams, opts ...grpc.CallOption) (*ContainerCreateResp, error)
	Start(ctx context.Context, in *ContainerStartParams, opts ...grpc.CallOption) (*ContainerStartResp, error)
	Stop(ctx context.Context, in *ContainerStopParams, opts ...grpc.CallOption) (*ContainerStopResp, error)
	Restart(ctx context.Context, in *ContainerRestartParams, opts ...grpc.CallOption) (*ContainerRestartResp, error)
}

type containerClient struct {
	cc grpc.ClientConnInterface
}

func NewContainerClient(cc grpc.ClientConnInterface) ContainerClient {
	return &containerClient{cc}
}

func (c *containerClient) Create(ctx context.Context, in *ContainerCreateParams, opts ...grpc.CallOption) (*ContainerCreateResp, error) {
	out := new(ContainerCreateResp)
	err := c.cc.Invoke(ctx, "/Container/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerClient) Start(ctx context.Context, in *ContainerStartParams, opts ...grpc.CallOption) (*ContainerStartResp, error) {
	out := new(ContainerStartResp)
	err := c.cc.Invoke(ctx, "/Container/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerClient) Stop(ctx context.Context, in *ContainerStopParams, opts ...grpc.CallOption) (*ContainerStopResp, error) {
	out := new(ContainerStopResp)
	err := c.cc.Invoke(ctx, "/Container/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerClient) Restart(ctx context.Context, in *ContainerRestartParams, opts ...grpc.CallOption) (*ContainerRestartResp, error) {
	out := new(ContainerRestartResp)
	err := c.cc.Invoke(ctx, "/Container/Restart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContainerServer is the server API for Container service.
// All implementations must embed UnimplementedContainerServer
// for forward compatibility
type ContainerServer interface {
	Create(context.Context, *ContainerCreateParams) (*ContainerCreateResp, error)
	Start(context.Context, *ContainerStartParams) (*ContainerStartResp, error)
	Stop(context.Context, *ContainerStopParams) (*ContainerStopResp, error)
	Restart(context.Context, *ContainerRestartParams) (*ContainerRestartResp, error)
	mustEmbedUnimplementedContainerServer()
}

// UnimplementedContainerServer must be embedded to have forward compatible implementations.
type UnimplementedContainerServer struct {
}

func (UnimplementedContainerServer) Create(context.Context, *ContainerCreateParams) (*ContainerCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedContainerServer) Start(context.Context, *ContainerStartParams) (*ContainerStartResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedContainerServer) Stop(context.Context, *ContainerStopParams) (*ContainerStopResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedContainerServer) Restart(context.Context, *ContainerRestartParams) (*ContainerRestartResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedContainerServer) mustEmbedUnimplementedContainerServer() {}

// UnsafeContainerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContainerServer will
// result in compilation errors.
type UnsafeContainerServer interface {
	mustEmbedUnimplementedContainerServer()
}

func RegisterContainerServer(s grpc.ServiceRegistrar, srv ContainerServer) {
	s.RegisterService(&Container_ServiceDesc, srv)
}

func _Container_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerCreateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Container/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServer).Create(ctx, req.(*ContainerCreateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Container_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerStartParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Container/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServer).Start(ctx, req.(*ContainerStartParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Container_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerStopParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Container/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServer).Stop(ctx, req.(*ContainerStopParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Container_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerRestartParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Container/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServer).Restart(ctx, req.(*ContainerRestartParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Container_ServiceDesc is the grpc.ServiceDesc for Container service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Container_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Container",
	HandlerType: (*ContainerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Container_Create_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Container_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Container_Stop_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _Container_Restart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "container_api.proto",
}
