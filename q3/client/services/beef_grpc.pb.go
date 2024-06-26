// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: beef.proto

package services

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
	Beef_CountBeef_FullMethodName = "/services.Beef/CountBeef"
)

// BeefClient is the client API for Beef service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BeefClient interface {
	CountBeef(ctx context.Context, in *BeefRequest, opts ...grpc.CallOption) (*BeefResponse, error)
}

type beefClient struct {
	cc grpc.ClientConnInterface
}

func NewBeefClient(cc grpc.ClientConnInterface) BeefClient {
	return &beefClient{cc}
}

func (c *beefClient) CountBeef(ctx context.Context, in *BeefRequest, opts ...grpc.CallOption) (*BeefResponse, error) {
	out := new(BeefResponse)
	err := c.cc.Invoke(ctx, Beef_CountBeef_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BeefServer is the server API for Beef service.
// All implementations must embed UnimplementedBeefServer
// for forward compatibility
type BeefServer interface {
	CountBeef(context.Context, *BeefRequest) (*BeefResponse, error)
	mustEmbedUnimplementedBeefServer()
}

// UnimplementedBeefServer must be embedded to have forward compatible implementations.
type UnimplementedBeefServer struct {
}

func (UnimplementedBeefServer) CountBeef(context.Context, *BeefRequest) (*BeefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountBeef not implemented")
}
func (UnimplementedBeefServer) mustEmbedUnimplementedBeefServer() {}

// UnsafeBeefServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BeefServer will
// result in compilation errors.
type UnsafeBeefServer interface {
	mustEmbedUnimplementedBeefServer()
}

func RegisterBeefServer(s grpc.ServiceRegistrar, srv BeefServer) {
	s.RegisterService(&Beef_ServiceDesc, srv)
}

func _Beef_CountBeef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeefServer).CountBeef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Beef_CountBeef_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeefServer).CountBeef(ctx, req.(*BeefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Beef_ServiceDesc is the grpc.ServiceDesc for Beef service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Beef_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.Beef",
	HandlerType: (*BeefServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CountBeef",
			Handler:    _Beef_CountBeef_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "beef.proto",
}
